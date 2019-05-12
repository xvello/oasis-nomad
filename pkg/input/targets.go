package input

import (
	log "github.com/sirupsen/logrus"
	billy "gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/util"
)

// Targets allows to specify a list of target files
// on an arbitrary filesystem
type Targets struct {
	filesystem billy.Filesystem
	selectors  []string
}

// Files returns all files scoped by the object via a channel
// User has to close files after usage
// Channel is closed when all files have been sent
func (t *Targets) Files() <-chan billy.File {
	output := make(chan billy.File)
	go func() {
		defer close(output)
		for _, s := range t.selectors {
			filenames, err := util.Glob(t.filesystem, s)
			if err != nil {
				log.WithFields(log.Fields{
					"selector": s,
					"error":    err,
				}).Warn("Invalid glob pattern ignored")
				continue
			}
			for _, name := range filenames {
				f, err := t.filesystem.Open(name)
				if err != nil {
					log.WithFields(log.Fields{
						"file":  name,
						"error": err,
					}).Warn("Cannot open file")
					continue
				}
				output <- f
			}
		}
	}()
	return output
}
