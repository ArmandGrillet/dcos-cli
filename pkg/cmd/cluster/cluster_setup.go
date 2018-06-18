package cluster

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-cli/pkg/setup"

	"github.com/spf13/cobra"
)

// newCmdClusterSetup configures the CLI with a given DC/OS cluster.
func newCmdClusterSetup(ctx api.Context) *cobra.Command {
	setupFlags := setup.NewFlags(ctx.Fs(), ctx.EnvLookup)
	cmd := &cobra.Command{
		Use:  "setup",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clusterURL := args[0]

			// We want to attach the cluster once it is setup.
			attach := true
			_, err := ctx.Setup(setupFlags, clusterURL, attach)
			return err
		},
	}
	setupFlags.Register(cmd.Flags())
	return cmd
}
