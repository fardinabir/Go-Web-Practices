package controllers

import (
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
)

func TakeBackup(args []string) {
	from := viper.GetString("backup.from")
	to := viper.GetString("backup.to")
	key := viper.GetString("backup.key")
	fmt.Println(from, to)

	cmdString := fmt.Sprintf("echo %s | restic -r %s backup %s", key, to, from)
	fmt.Println(cmdString)
	execCmd(cmdString)
}

func InitRepo(args []string) {
	config := GetFromConfig()

	for _, val := range args {
		location, ok := config.Locations[val]
		if !ok {
			fmt.Println(fmt.Errorf("given location not found in config").Error())
			return
		}
		fmt.Println(location)
		backendNames := location.To

		var backendList []Backend
		for _, backendName := range backendNames {
			if backend, ok := config.Backends[backendName]; ok {
				fmt.Println(backend)
				backendList = append(backendList, backend)
			}
		}

		newRepo := Repo{
			Name:        val,
			From:        location.From,
			Backends:    backendList,
			Initialized: false,
		}
		newRepo.Init()
	}
}

func Snapshots(args []string) {
	repo := viper.GetString("snapshots.path")
	key := viper.GetString("snapshots.key")
	fmt.Println(repo, key)

	cmdString := fmt.Sprintf("echo %s | restic snapshots -r %s", key, repo)
	fmt.Println(cmdString)
	execCmd(cmdString)
}

func Restore(args []string) {
	from := viper.GetString("restore.from")
	to := viper.GetString("restore.to")
	key := viper.GetString("restore.key")
	snapshotId := viper.GetString("restore.snapshot")
	if snapshotId == "" {
		snapshotId = "latest"
	}
	cmdString := fmt.Sprintf("echo %s | restic restore -r %s --target %s %s", key, from, to, snapshotId)
	fmt.Println(cmdString)
	execCmd(cmdString)
}

func execCmd(cmdString string) {
	cmd := exec.Command("bash", "-c", cmdString)

	// CombinedOutput() runs the command and returns its combined standard output and standard error.
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Command output:", string(output))
}
