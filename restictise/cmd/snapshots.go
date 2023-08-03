package cmd

import (
	"github.com/spf13/cobra"
	"restictise/controllers"
)

// snapshotsCmd represents the snapshots command
var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "To view the taken snapshots",
	Long:  `To view the taken snapshots`,
	Run: func(cmd *cobra.Command, args []string) {
		controllers.Snapshots(args)
	},
}

func init() {
	rootCmd.AddCommand(snapshotsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapshotsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapshotsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
