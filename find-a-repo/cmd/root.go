/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "find-a-rep",
	Short: "Simple cli tool for finding related github repo",
	Long:  `find-a-repo is a lightweight cli tool for finding a suitable repository on the given topic/tool name.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: runSearch,
}

//func runSearch(cmd *cobra.Command, args []string) {
//	finders.FindRepos(args)
//}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.find-a-repo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	println("Init of root...............")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
