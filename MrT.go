package main

import (
	"github.com/kcmerrill/MrT/cmd"
	"github.com/spf13/viper"
	"os"
	"os/user"
)

func main() {
	/* Set basic defaults */
	viper.SetDefault("editor", "vim")
	viper.SetDefault("editor_args", "+startinsert -c \"normal O\"")
	viper.SetDefault("project", "personal")
	viper.SetDefault("date_format", "1/2/06@3:04pm")

	/* Set the default storage */
	if u, err := user.Current(); err == nil {
		viper.SetDefault("tasks", u.HomeDir+"/MrT/tasks")
		viper.SetDefault("tasks_backup", u.HomeDir+"/MrT/.tasks.bkup")
	}

	/* Is there a tasks file in the current directory? If so, lets use it */
	if cur_dir, e := os.Getwd(); e == nil {
		if _, err := os.Stat(cur_dir + "/tasks"); err == nil {
			viper.SetDefault("tasks", cur_dir+"/tasks")
			viper.SetDefault("tasks_backup", cur_dir+"/.tasks.bkup")
		}
	}

	/* Giddy Up! */
	cmd.Execute()
}
