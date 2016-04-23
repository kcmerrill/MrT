package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new tasks list in the current folder you are in",
	Run: func(cmd *cobra.Command, args []string) {
		entries.Init()
		display.Init()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
