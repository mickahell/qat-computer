package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	ConfPath         string
	LogLevel         int      `yaml:"loglevel"`
	ComputePath      string   `yaml:"compute_path"`
	FileExeName      string   `yaml:"filename_to_execute"`
	RequirementsFile string   `yaml:"requirements_file"`
	PythonVer        string   `yaml:"python_version"`
	DebianPkg        []string `yaml:"debian_packages"`
}

var AppConfig Cfg
var err error

// flags
var ConfPathflag string
var LogLevelflag int
var ComputePathflag string
var FileExeNameflag string
var RequirementsFileflag string
var PythonVerflag string
var DebianPkgflag string

func InitFile() {
	AppConfig.LogLevel = LogLevelflag
	if filepath.IsAbs(ComputePathflag) {
		AppConfig.ComputePath = ComputePathflag
	} else {
		AppConfig.ComputePath, err = filepath.Abs(filepath.Join("compute", ComputePathflag))
		if err != nil {
			fmt.Println(err)
		}
	}
	AppConfig.FileExeName = FileExeNameflag
	AppConfig.RequirementsFile = RequirementsFileflag
	AppConfig.PythonVer = PythonVerflag
	if DebianPkgflag != "" {
		AppConfig.DebianPkg = strings.Split(DebianPkgflag, " ")
	}

	if ConfPathflag != "" {
		if filepath.IsAbs(ConfPathflag) {
			AppConfig.ConfPath = ConfPathflag
		} else {
			AppConfig.ConfPath, err = filepath.Abs(filepath.Join("conf", ConfPathflag))
			if err != nil {
				fmt.Println(err)
			}
		}
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

	if !filepath.IsAbs(AppConfig.ComputePath) {
		AppConfig.ComputePath, err = filepath.Abs(filepath.Join("compute", AppConfig.ComputePath))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TheAppConfig() *Cfg {
	return &AppConfig
}
