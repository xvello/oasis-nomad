package commands

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/xvello/oasis-nomad/pkg/nomad"
)

var (
	waitCmd = &cobra.Command{
		Use:   "wait",
		Short: "Wait until nomad is available",
		Args:  cobra.ExactArgs(0),
		RunE:  wait,
	}
)

func init() {
	OasisCmd.AddCommand(waitCmd)
	waitCmd.Flags().DurationP("frequency", "f", time.Second, "time between two retries")
	waitCmd.Flags().DurationP("timeout", "t", 10*time.Second, "maximum time before aborting")
}

func wait(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	f, err := cmd.Flags().GetDuration("frequency")
	if err != nil {
		return err
	}
	t, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		return err
	}
	return cli.Wait(f, t)
}
