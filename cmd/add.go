package cmd

import (
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/editor"
	"github.com/kcmerrill/MrT/entries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a task/note",
	Aliases: []string{"new"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			new_task := strings.Join(args, " ")
			entries.New(new_task)
			entries.Update()
		} else {
			editor.Run(viper.GetString("editor"), viper.GetString("editor_args"), viper.GetString("tasks"))
			entries.Update()
		}
		display.Added()
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
