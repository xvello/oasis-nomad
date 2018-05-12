package nomad

import (
	"errors"

	"github.com/hashicorp/nomad/api"
	log "github.com/sirupsen/logrus"

	"github.com/xvello/ndeploy/pkg/docker/registry"
)

func (c *Client) updateAndRun(job *api.Job) error {
	err := addDigests(job)
	if err != nil {
		return err
	}

	response, _, err := c.cli.Jobs().Register(job, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"job":   jobName(job),
			"error": err,
		}).Warn("Cannot register job")
		return err
	}

	if response.Warnings != "" {
		log.WithFields(log.Fields{
			"job":      jobName(job),
			"warnings": response.Warnings,
			"evalID":   response.EvalID,
		}).Warning("Registered job with warnings")
	} else {
		log.WithFields(log.Fields{
			"job":    jobName(job),
			"evalID": response.EvalID,
		}).Debug("Registered job OK")
	}

	return nil
}

// addDigests adds docker image digests as
// in the job spec, to force auto-updates
func addDigests(job *api.Job) error {
	if job == nil {
		return errors.New("nil job pointer")
	}

	for _, group := range job.TaskGroups {
		if group == nil {
			continue
		}
		for _, task := range group.Tasks {
			if task == nil {
				continue
			}
			if task.Driver != "docker" {
				continue
			}
			image, found := task.Config["image"]
			if !found {
				continue
			}
			imgString, ok := image.(string)
			if !ok {
				log.WithFields(log.Fields{
					"job":   jobName(job),
					"group": groupName(group),
					"task":  task.Name,
				}).Warning("Invalid image field")
				continue
			}

			digest, err := registry.ResolveFromString(imgString)
			if err != nil {
				log.WithFields(log.Fields{
					"job":   jobName(job),
					"group": groupName(group),
					"task":  task.Name,
					"image": imgString,
					"error": err,
				}).Warning("Cannot get digest, skiping")
				continue
			}

			task.Config["image"] = digest.String()
		}
	}

	return nil
}
