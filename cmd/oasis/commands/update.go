package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/nomad"
)

var (
	updateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"u"},
		Short:   "Update all running jobs",
		Args:    cobra.ExactArgs(0),
		RunE:    update,
	}
)

func init() {
	OasisCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntP("parallelism", "p", 4, "updates to run in parallel")
}

func update(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	p, err := cmd.Flags().GetInt("parallelism")
	if err != nil {
		return err
	}
	if p < 1 {
		return fmt.Errorf("Parallelism must be higher than 0: got %d", p)
	}

	return cli.Update(p)
}
