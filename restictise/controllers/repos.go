package controllers

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	Name        string
	From        string
	Backends    []Backend
	Initialized bool
}

func (r Repo) Backup() error {
	cmdString := fmt.Sprintf("echo %s | restic -r %s backup %s", r.Backends[0].Key, r.Backends[0].Path, r.From)
	fmt.Println(cmdString)
	execCmd(cmdString)
	return nil
}

func (r Repo) Init() error {
	if r.Initialized {
		fmt.Println("Already Initialized !")
		return errors.New("Already Initialized")
	}
	for _, backend := range r.Backends {
		if backend.Env != nil {
			for key, value := range backend.Env {
				os.Setenv(strings.ToUpper(key), value)
			}
		}
		cmdString := fmt.Sprintf("echo %s | echo %s | restic -r %s init", backend.Key, backend.Key, backend.Path)
		fmt.Println(cmdString, "--------------------")
		execCmd(cmdString)
	}

	r.Initialized = true
	return nil
}
