package main

import (
	"github.com/kcmerrill/MrT/cmd"
	"github.com/spf13/viper"
	"os/user"
)

func main() {
	/* Set basic defaults */
	viper.SetDefault("editor", "vim")
	viper.SetDefault("editor_args", "+startinsert -c \"normal O\"")
	viper.SetDefault("project", "personal")
	viper.SetDefault("date_format", "1/2/06@3:04")

	/* Set the default storage */
	if u, err := user.Current(); err == nil {
		viper.SetDefault("tasks", u.HomeDir+"/MrT/tasks")
		viper.SetDefault("tasks_backup", u.HomeDir+"/MrT/.tasks.bkup")
	} else {
		/* TODO: Add an error message */
	}

	/* Giddy Up! */
	cmd.Execute()
}
