package main

import "github.com/spf13/cobra"

func GetRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "executor",
		Long: "you can set the api path by `export EXECUTOR_API=unix|tcp://path`",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
