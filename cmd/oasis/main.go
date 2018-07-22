package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/xvello/oasis-nomad/cmd/oasis/commands"
)

func main() {
	if err := commands.OasisCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
