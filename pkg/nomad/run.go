package nomad

import (
	"fmt"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/jobspec"
	log "github.com/sirupsen/logrus"
	"github.com/xvello/oasis-nomad/pkg/input"
)

// Run parses job specs, resolves the image digests and runs the jobs
func (c *Client) Run(targets *input.Targets) error {
	var total, hasError int

	for f := range targets.Files() {
		total++
		job, err := jobspec.Parse(f)
		f.Close()
		if err != nil {
			log.WithFields(log.Fields{
				"file":  f.Name(),
				"error": err,
			}).Warn("Cannot parse file")
			hasError++
			continue
		}
		err = c.updateAndRun(job, true)
		if err != nil {
			hasError++
			continue
		}
	}

	if hasError > 0 {
		return fmt.Errorf("%d out of %d jobs errored out", hasError, total)
	}
	return nil
}

// Update updates all running jobs on the Nomad cluster to the latest digest
func (c *Client) Update(parallelism int) error {
	hasError := 0

	jobs, _, err := c.cli.Jobs().List(nil)
	if err != nil {
		return err
	}

	// Start workers
	jobChan := make(chan *api.JobListStub)
	defer close(jobChan)
	errChan := make(chan error)
	defer close(errChan)
	for i := 0; i < parallelism; i++ {
		go c.jobSubmitRoutine(jobChan, errChan)
	}

	// Submit jobs
	go func() {
		for _, j := range jobs {
			jobChan <- j
		}
	}()

	// Wait for results
	var finished int
	for err := range errChan {
		if err != nil {
			hasError++
		}
		finished++
		if finished == len(jobs) {
			break
		}
	}

	if hasError > 0 {
		return fmt.Errorf("%d out of %d jobs errored out", hasError, len(jobs))
	}
	return nil
}

func (c *Client) jobSubmitRoutine(jobs <-chan *api.JobListStub, errs chan<- error) {
	for j := range jobs {
		if j == nil {
			errs <- nil
			continue
		}
		job, _, err := c.cli.Jobs().Info(j.ID, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"job":   jobName(job),
				"error": err,
			}).Warn("Cannot retrieve job")
			errs <- err
			continue
		}
		err = c.updateAndRun(job, true)
		errs <- err
	}
}

// Reset deletes all running jobs on the Nomad cluster
func (c *Client) Reset() error {
	hasError := 0

	jobs, _, err := c.cli.Jobs().List(nil)
	if err != nil {
		return err
	}

	for _, j := range jobs {
		if j == nil {
			continue
		}
		_, _, err := c.cli.Jobs().Deregister(j.ID, true, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"job":   j.Name,
				"error": err,
			}).Warn("Cannot deregister job")
			hasError++
			continue
		}
	}

	if hasError > 0 {
		return fmt.Errorf("%d out of %d jobs errored out", hasError, len(jobs))
	}
	return nil
}
