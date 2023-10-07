package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	ConfPath    string
	LogLevel    int    `yaml:"loglevel"`
	ComputePath string `yaml:"computepath"`
}

var AppConfig Cfg

// flags
var ConfPathflag string
var ComputePathflag string
var LogLevelflag int

func InitFile() {
	AppConfig.LogLevel = LogLevelflag
	AppConfig.ComputePath = ComputePathflag

	if ConfPathflag != "" {
		AppConfig.ConfPath = ConfPathflag
		ReadConfig()
	}
}

func ReadConfig() {
	file, err := os.Open(AppConfig.ConfPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		fmt.Println(err)
	}
}

func TheAppConfig() *Cfg {
	return &AppConfig
}
