package controllers

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

type Config struct {
	Locations map[string]Location
	Backends  map[string]Backend
}

type Location struct {
	Type string
	Name string
	From string
	To   []string
}

type Backend struct {
	Type string
	Name string
	Path string
	Key  string
	Env  map[string]string
}

var config *Config

func GetFromConfig() *Config {
	if config == nil {
		once.Do(func() {
			config = &Config{}
			if err := viper.UnmarshalExact(config); err != nil {
				fmt.Println("Couldn't parse config file")
			}
			fmt.Println(config.Backends["hdd"])
		})
	}
	return config
}

//func GetBackends() *[]Backends {
//	//var backendArr *[]Backends
//	//vv := viper.Get("backends")
//	//return *vv
//}
