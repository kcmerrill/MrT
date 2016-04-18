package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
	"strconv"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Complete a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()
		task_id := 0
		if len(args) >= 1 {
			if ti, err := strconv.Atoi(args[0]); err == nil {
				task_id = ti
			} else {
				display.Error("Task # must be a number.")
			}
		}
		if e, err := entries.Complete(task_id); err == nil {
			display.Complete(task_id, e)
		} else {
			display.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
