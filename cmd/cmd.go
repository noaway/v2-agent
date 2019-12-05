package cmd

import (
	"github.com/spf13/cobra"
)

var configPath string

func Commands(root *cobra.Command, childs ...*cobra.Command) {
	root.AddCommand(
		webCommand(),
		agentCommand(),
		conversionCommand(),
		testCommand(),
	)
	root.Flags().StringVarP(&configPath, "config", "c", "", "config path")
}
