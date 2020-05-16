package image

import (
	"github.com/spf13/cobra"
)

func GetImageCmd() *cobra.Command {
	imageCmd := &cobra.Command{
		Use:   "image [list/delete/import]",
		Short: "operations of images",
		Long:  `the operations of the images for example:list, import, delete`,
		Args:  cobra.MinimumNArgs(0),
		Run:   func(cmd *cobra.Command, args []string) { _ = cmd.Help() },
	}
	imageCmd.AddCommand(
		GetImageListCmd(),
		GetImageImportCmd(),
		GetImageDeleteCmd(),
		GetImageExportCmd(),
	)
	return imageCmd
}

func handle(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

//func CheckNameOrId(imageId string) string {
//	imageListResp, err := connection.Client.ListImage(context.Background(), &empty.Empty{})
//	if err != nil {
//		panic(err)
//	}
//
//	isNameFlag := false
//	isIdFlag := false
//
//	for _, v := range imageListResp.Images {
//		if imageId == v.Name {
//			isNameFlag = true
//			imageId = v.Id
//			break
//		}
//
//		if len(imageId) == 12 {
//			if imageId == v.Id[:12] {
//				isIdFlag = true
//				imageId = v.Id
//				break
//			}
//		} else {
//			if imageId == v.Id {
//				isIdFlag = true
//				imageId = v.Id
//				break
//			}
//		}
//	}
//
//	if !(isNameFlag || isIdFlag) {
//		panic("No Image named: " + imageId)
//	}
//	return imageId
//}
