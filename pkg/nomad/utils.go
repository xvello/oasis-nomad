package nomad

import (
	"path/filepath"

	"github.com/hashicorp/nomad/api"
)

func jobName(j *api.Job) string {
	if j == nil {
		return "nil"
	}
	if j.Name != nil {
		return *j.Name
	}
	if j.ID != nil {
		return *j.ID
	}
	return "unknown"
}

func groupName(g *api.TaskGroup) string {
	if g == nil {
		return "nil"
	}
	if g.Name != nil {
		return *g.Name
	}
	return "unknown"
}

func globFileList(args []string) []string {
	var files []string
	for _, pattern := range args {
		f, _ := filepath.Glob(pattern)
		files = append(files, f...)
	}
	return files
}
