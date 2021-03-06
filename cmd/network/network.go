package network

import "github.com/spf13/cobra"

func GetNetworkCmd() *cobra.Command {
	networkCmd := &cobra.Command{
		Use:   "network [create/inspect/delete/list] [options]",
		Short: "operations of networks",
		Long:  `the operations of the networks for example:create, delete`,
		Args:  cobra.MinimumNArgs(0),
		Run:   handle,
	}
	networkCmd.AddCommand(
		GetNetworkCreateCmd(),
		GetNetworkDeleteCmd(),
		GetNetworkInspectCmd(),
		GetNetworkListCmd(),
		GetNetworkConnectCmd(),
	)
	return networkCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
