package plugin

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/spf13/cobra"
)

// NewCommand creates the `dcos auth` subcommand.
func NewCommand(ctx api.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Manage CLI plugins",
	}
	cmd.AddCommand(
		newCmdPluginAdd(ctx),
		newCmdPluginRemove(ctx),
	)
	return cmd
}
