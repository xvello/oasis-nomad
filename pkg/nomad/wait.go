package nomad

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// Wait allows to block until the nomad server is healthy
func (c *Client) Wait(freq, max time.Duration) error {
	ticker := time.NewTicker(freq)
	defer ticker.Stop()
	timeout := time.NewTimer(max)
	defer timeout.Stop()

	var err error
	for {
		select {
		case <-ticker.C:
			err = c.serverHealthy()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Debug("Nomad server unhealthy")
				continue
			}
			log.Info("Nomad server healthy")
			return nil
		case <-timeout.C:
			return err
		}
	}
}

func (c *Client) serverHealthy() error {
	health, err := c.cli.Agent().Health()
	if err != nil {
		return err
	}
	if health.Server == nil {
		return fmt.Errorf("invalid health response")
	}
	if !health.Server.Ok {
		return fmt.Errorf("unhealthy server: %s", health.Server.Message)
	}
	return nil
}
