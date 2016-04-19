package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
	"strconv"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a task",
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
		if e, err := entries.Start(task_id); err == nil {
			display.Start(task_id, e)
		} else {
			display.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
