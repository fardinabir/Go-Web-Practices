/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "restictise",
	Short: "To work with restic in more flexible ways",
	Long:  `To work with restic in more flexible ways`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func loadConfig() {
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	//viper.SetConfigFile("v1..config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, ", err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	loadConfig()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
