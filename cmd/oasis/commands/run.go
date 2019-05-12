package commands

import (
	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/input"
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
	runCmd.Flags().StringP("git-url", "g", "", "url of the git repository to checkout")

}

func run(cmd *cobra.Command, args []string) error {
	gitURL, err := cmd.Flags().GetString("git-url")
	if err != nil {
		return err
	}
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}

	var targets *input.Targets
	if len(gitURL) > 0 {
		targets, err = input.FromGit(gitURL, args)
	} else {
		targets, err = input.Local(args)
	}
	if err != nil {
		return err
	}
	return cli.Run(targets)
}
