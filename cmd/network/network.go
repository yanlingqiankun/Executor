package network

import "github.com/spf13/cobra"

func GetNetworkCmd() *cobra.Command {
	networkCmd := &cobra.Command{
		Use:   "network [create/inspect/remove/rename/list]",
		Short: "operations of networks",
		Long:  `the operations of the networks for example:create, remove`,
		Args:  cobra.MinimumNArgs(0),
		Run:   handle,
	}
	networkCmd.AddCommand(
		GetNetworkCreateCmd(),
		GetNetworkDeleteCmd(),
		//GetNetworkRemoveCmd(),
		//GetNetworkInspectCmd(),
		//GetNetworkListCmd(),
		//GetNetworkRecoveryCmd(),
	)
	return networkCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
