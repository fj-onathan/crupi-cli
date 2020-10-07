package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Viper config location
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "crupi",
	Short: "CRUD API creator with custom dummy data.",
	Long:  `CRUD API creator and server, with custom dummy data using CLI mode in Go.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

// Execute will execute the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".config.json", "config file location (defaults to .config.json)")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Config file not found at path %s\n", cfgFile)
		} else {
			panic(fmt.Errorf("Error: uncaught error! %s", err))
		}
	}
}
