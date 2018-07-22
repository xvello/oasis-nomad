package commands

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// OasisCmd is the root command
	OasisCmd = &cobra.Command{
		Use:          "oasis [command]",
		Short:        "Deploys and reports on Nomad jobs",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := setupLogger(); err != nil {
				log.Error(err)
				os.Exit(-1)
			}
		},
	}

	logLevel string
)

func init() {
	OasisCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "info", "logging level")
}

func setupLogger() error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	return nil
}
