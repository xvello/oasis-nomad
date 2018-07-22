package main

import (
	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/nomad"
)

func run(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	return cli.Run(args)
}

func reset(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	return cli.Reset()
}

func update(cmd *cobra.Command, args []string) error {
	cli, err := nomad.Connect(nil)
	if err != nil {
		return err
	}
	return cli.Update()
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
