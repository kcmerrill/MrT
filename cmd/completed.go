package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/kcmerrill/MrT/entry"
	"github.com/spf13/cobra"
	"time"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows a list of completed tasks within a given duration",
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()
		dur, _ := time.ParseDuration("7h")
		if len(args) >= 1 {
			if d, err := time.ParseDuration(args[0]); err == nil {
				dur = d
			}
		}
		entries.List(-1, func(e *entry.Entry) bool {
			if completed, err := e.Completed(); err == nil {
				if completed.After(time.Now().Add(-dur)) {
					return true
				} else {
					return false
				}
			}
			return false
		})
		display.LS()
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
