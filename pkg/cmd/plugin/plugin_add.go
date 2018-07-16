package plugin

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/spf13/cobra"
)

// newCmdPluginAdd creates the `dcos plugin add` subcommand.
func newCmdPluginAdd(ctx api.Context) *cobra.Command {
	var name string
	var type string
	var update bool
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a CLI plugin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, err := ctx.Cluster()
			if err != nil {
				return err
			}

			pluginManager := ctx.PluginManager(cluster.SubcommandsDir())

			pluginManager.Install()

			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Name of the plugin")
	cmd.Flags().StringVar(&type, "type", plugin.TypeZip, `Type of the plugin ("zip" or "executable")`)
	cmd.Flags().BoolVarP(&update, "update", false, "Update an already existing plugin")
	return cmd
}
