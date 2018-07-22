package commands

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/xvello/oasis-nomad/pkg/releases"
)

var (
	agentCmd = &cobra.Command{
		Use:   "agent [command]",
		Short: "Manage the Nomad agent binary",
	}
	agentUpgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the agent",
		Args:  cobra.ExactArgs(0),
		RunE:  agentUpgrade,
	}
	installPrefix string
)

func init() {
	OasisCmd.AddCommand(agentCmd)
	agentCmd.PersistentFlags().StringVarP(&installPrefix, "prefix", "p", "/usr/local", "installation prefix")
	agentCmd.AddCommand(agentUpgradeCmd)
	agentUpgradeCmd.Flags().StringP("os", "o", "linux", "target OS")
	agentUpgradeCmd.Flags().StringP("arch", "a", "amd64", "target architecture")
}

func agentUpgrade(cmd *cobra.Command, args []string) error {
	buildOs, err := cmd.Flags().GetString("os")
	if err != nil {
		return err
	}
	buildArch, err := cmd.Flags().GetString("arch")
	if err != nil {
		return err
	}
	prefix, err := cmd.Flags().GetString("prefix")
	if err != nil {
		return err
	}

	nomadVersions, err := releases.GetReleases("nomad")
	if err != nil {
		return err
	}

	latest, err := nomadVersions.Latest(true)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"version": latest.Version,
	}).Info("Found nomad version")

	dest := filepath.Join(prefix, "bin", "nomad")
	destVer := fmt.Sprintf("%s-%s", dest, latest.Version)

	_ = os.Remove(destVer)
	err = latest.Download(destVer, buildOs, buildArch)
	if err != nil {
		return err
	}
	_ = os.Remove(dest)
	err = os.Symlink(destVer, dest)
	if err != nil {
		return err
	}

	return nil
}
