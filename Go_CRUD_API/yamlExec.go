package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func readFromYaml() {
	yamlFile, err := os.Open("./dbconfig.yaml")
	if err != nil {
		log.Println(err)
	}
	defer yamlFile.Close()

	decoder := yaml.NewDecoder(yamlFile)
	var config DBConfig
	if err := decoder.Decode(&config); err != nil {
		log.Println(err)
	}
	fmt.Println(config.Host, config.User, config.DBName)
}
