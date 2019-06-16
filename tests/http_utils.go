package tests

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hashicorp/nomad/api"
)

var httpClient = &http.Client{Timeout: 2 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func listNomadJobs() ([]*api.JobListStub, error) {
	jobs := []*api.JobListStub{}
	err := getJSON(nomadBaseURL+"/v1/jobs", &jobs)
	return jobs, err
}

func getNomadJob(name string) (*api.Job, error) {
	job := new(api.Job)
	err := getJSON(nomadBaseURL+"/v1/job/"+name, &job)
	return job, err
}
