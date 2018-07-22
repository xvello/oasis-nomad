package commands

import (
	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/nomad"
)

var (
	runCmd = &cobra.Command{
		Use:     "run [files]",
		Aliases: []string{"r"},
		Short:   "Runs nomad jobs from files",
		Args:    cobra.MinimumNArgs(1),
		RunE:    run,
	}
)

func init() {
	OasisCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	return cli.Run(args)
}
