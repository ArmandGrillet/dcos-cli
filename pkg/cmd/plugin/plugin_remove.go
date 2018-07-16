package plugin

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-cli/pkg/cluster/linker"
	"github.com/dcos/dcos-cli/pkg/config"
	"github.com/spf13/cobra"
)

// newCmdPluginRemove creates the `dcos plugin remove` command.
func newCmdPluginRemove(ctx api.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a CLI plugin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			attachedCluster, err := ctx.Cluster()
			if err != nil {
				return err
			}

			manager := ctx.ConfigManager()
			linkedClusterConfig, err := manager.Find(args[0], false)
			if err != nil {
				return err
			}
			linkedCluster := config.NewCluster(linkedClusterConfig)

			attachedClient := linker.New(ctx.HTTPClient(attachedCluster), ctx.Logger())
			return attachedClient.Unlink(linkedCluster.ID())
		},
	}
	return cmd
}
