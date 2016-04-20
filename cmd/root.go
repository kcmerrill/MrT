package cmd

import (
	"fmt"
	"github.com/kcmerrill/MrT/display"
	"github.com/kcmerrill/MrT/entries"
	"github.com/kcmerrill/MrT/entry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "MrT",
	Short: "A command line task manager",
	Run: func(cmd *cobra.Command, args []string) {
		entries.Update()
		entries.List(1, func(e *entry.Entry) bool {
			return !e.IsCompleted()
		})
		display.Current()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/MrT/config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config")     // name of config file (without extension)
	viper.AddConfigPath("$HOME/MrT/") // adding home directory as first search path
	viper.AutomaticEnv()              // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
	}
}
