/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go_CRUD_API/server"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "A go crud api server",
	Long:  `A go crud api server that allows login, registration with basic functionalities`,
	Run:   runServ,
}

func runServ(cmd *cobra.Command, args []string) {
	fmt.Println("This is root cmd")
	loadFromConfig(cmd)
	fmt.Println("this is pFlag-Viper", viper.GetString("port"))
	fmt.Println("this is pFlag-Cobra", cmd.Flags().Lookup("port"))
	portNow, _ := cmd.Flags().GetString("port")
	server.StartServer(portNow)
}

func loadFromConfig(cmd *cobra.Command) {
	// For reading flag values from config file, which were not given during cmd input
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			val := viper.GetString("server." + f.Name) // as they are under sever tag in config file
			if val != "" {                             // checks if config file contains value for the flag
				fmt.Println("This is from config : ", f.Name, " ", val)
				cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			}
		}
	})
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("port", "p", "3333", "Test string cmd")

	// init viper for reading config file
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, ", err)
	}

	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
}
