package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Display a list of tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()
		entries.List(10)
		display.LS()
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
