/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restictise/controllers"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup using restic",
	Long:  `Backup using restic`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.TakeBackup(args)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
