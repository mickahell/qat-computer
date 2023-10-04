package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	Loglevel int `yaml:"loglevel"`
	ConfPath string
}

var AppConfig Cfg

// flags
var Confpathflag string
var LogLevelflag int

func InitFile() {
	AppConfig.Loglevel = LogLevelflag

	if Confpathflag != "" {
		AppConfig.ConfPath = Confpathflag
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
