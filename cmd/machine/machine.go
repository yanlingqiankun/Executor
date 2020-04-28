package machine

import "github.com/spf13/cobra"

func GetMachineCmd() *cobra.Command {
	machineCmd := &cobra.Command{
		Use:   "machine [create/inspect/start/stop/kill/restart/delete/rename/pause/unpause/attach]",
		Short: "operations of machines",
		Long:  `the operations of the machines for example:start, create, delete`,
		Args:  cobra.MinimumNArgs(0),
		Run:   handle,
	}
	machineCmd.AddCommand(
		GetMachineCreateCmd(),
		GetMachineKillCmd(),
		GetMachineDeleteCmd(),
		GetMachineRenameCmd(),
		GetMachineRestartCmd(),
		GetMachineStartCmd(),
		GetMachineStopCmd(),
		GetMachineInspectCmd(),
		GetMachineListCmd(),
		GetAttachCmd(),
		GetMachinePauseCmd(),
		GetMachineUnpauseCmd(),
	)
	return machineCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

