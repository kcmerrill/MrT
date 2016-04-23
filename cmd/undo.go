package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:     "undo",
	Short:   "Undo the last action you performed",
	Aliases: []string{"revert", "whoops"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := entries.Undo(); err == nil {
			display.Undo()
		} else {
			display.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(undoCmd)
}
