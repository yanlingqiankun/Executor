package machine

import "github.com/spf13/cobra"

func GetMachineCmd() *cobra.Command {
	machineCmd := &cobra.Command{
		Use:   "machine [create/inspect/start/stop/kill/restart/remove/rename]",
		Short: "operations of machines",
		Long:  `the operations of the machines for example:start, create, remove`,
		Args:  cobra.MinimumNArgs(0),
		Run:   handle,
	}
	machineCmd.AddCommand(
		GetMachineCreateCmd(),
		//GetContainerKillCmd(),
		//GetContainerRemoveCmd(),
		//GetContainerRenameCmd(),
		//GetContainerRestartCmd(),
		//GetContainerStartCmd(),
		//GetContainerStopCmd(),
		//GetContainerInspectCmd(),
		//GetContainerListCmd(),
		//GetContainerPauseCmd(),
		//GetContainerUnpauseCmd(),
		//GetContainerWaitCmd(),
	)
	return machineCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

