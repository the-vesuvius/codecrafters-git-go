package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mygit",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() error {

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(catFile)
	rootCmd.AddCommand(hashObject)
	rootCmd.AddCommand(lsTree)

	return rootCmd.Execute()
}
