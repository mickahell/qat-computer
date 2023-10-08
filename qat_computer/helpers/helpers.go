package helpers

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	ConfPath    string
	LogLevel    int      `yaml:"loglevel"`
	ComputePath string   `yaml:"computepath"`
	PythonVer   string   `yaml:"python_version"`
	DebianPkg   []string `yaml:"debian_packages"`
}

var AppConfig Cfg

// flags
var ConfPathflag string
var LogLevelflag int
var ComputePathflag string
var PythonVerflag string
var DebianPkgflag string

func InitFile() {
	AppConfig.LogLevel = LogLevelflag
	AppConfig.ComputePath = ComputePathflag
	AppConfig.PythonVer = PythonVerflag
	AppConfig.DebianPkg = strings.Split(DebianPkgflag, " ")

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
