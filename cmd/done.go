package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Complete a task",
	Aliases: []string{"complete", "finish"},
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()

		/* Did you pass in anything? No, default to 0 then */
		if len(args) == 0 {
			args = append(args, "0")
		}

		/* Loop through everything, mark tasks completed that exist */
		for _, id := range args {
			if ti, err := strconv.Atoi(id); err == nil {
				if e, err := entries.Complete(ti); err == nil {
					display.Complete(ti, e)
				} else {
					display.Error(err.Error())
				}
			} else {
				display.Error("Task # must be a number.")
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
