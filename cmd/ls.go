package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/kcmerrill/MrT/entry"
	"github.com/spf13/cobra"
)

var completed bool

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "Display a list of tasks",
	Aliases: []string{"list", "show"},
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()
		if completed {
			entries.List(10, nil)
			display.LS(
				[]string{"ID", "Description", "Created", "Completed"},
			)
		} else {
			entries.List(10, func(e *entry.Entry) bool {
				return !e.IsCompleted()
			})
			display.LS(nil)
		}
	},
}

func init() {
	lsCmd.Flags().BoolVarP(&completed, "all", "a", false, "Displays completed tasks")
	RootCmd.AddCommand(lsCmd)
}
