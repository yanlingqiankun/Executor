package volume

import (
	"github.com/spf13/cobra"
)

func GetVolumeCmd() *cobra.Command {
	volumeCmd := &cobra.Command{
		Use:   "volume [list/remove/create]",
		Short: "operations of volumes",
		Long:  `the operations of the volumes for example:list, create, remove`,
		Args:  cobra.MinimumNArgs(0),
		Run:   handle,
	}

	volumeCmd.AddCommand(
		GetVolumeAddCmd(),
		GetVolumeListCmd(),
		GetVolumeCreateCmd(),
		GetVolumeDeleteCmd(),
	)
	return volumeCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
