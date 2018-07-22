package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	oasisCmd = &cobra.Command{
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
	runCmd = &cobra.Command{
		Use:     "run [files]",
		Aliases: []string{"r"},
		Short:   "Runs nomad jobs from files",
		Args:    cobra.MinimumNArgs(1),
		RunE:    run,
	}
	updateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"u"},
		Short:   "Update all running jobs",
		Args:    cobra.ExactArgs(0),
		RunE:    update,
	}
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset all nomad jobs",
		Args:  cobra.ExactArgs(0),
		RunE:  reset,
	}
	waitCmd = &cobra.Command{
		Use:   "wait",
		Short: "Wait until nomad is available",
		Args:  cobra.ExactArgs(0),
		RunE:  wait,
	}
	logLevel string
)

func init() {
	oasisCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "info", "logging level")
	oasisCmd.AddCommand(runCmd)
	oasisCmd.AddCommand(resetCmd)
	oasisCmd.AddCommand(updateCmd)
	oasisCmd.AddCommand(waitCmd)
	waitCmd.Flags().DurationP("frequency", "f", time.Second, "time between two retries")
	waitCmd.Flags().DurationP("timeout", "t", 10*time.Second, "maximum time before aborting")
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

func main() {
	if err := oasisCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
