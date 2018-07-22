package commands

import (
	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/nomad"
)

var (
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset all nomad jobs",
		Args:  cobra.ExactArgs(0),
		RunE:  reset,
	}
)

func init() {
	OasisCmd.AddCommand(resetCmd)
}

func reset(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	return cli.Reset()
}
