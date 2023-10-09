package docs

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"qat-computer/logger"
)

//go:embed VERSION.txt
var version_file string
var app_name string = "Qat Computer"

const ComputeStr = "compute"
const draw = `
  #####                   #####
 #     #   ##   #####    #     #  ####  #    # #####  #    # ##### ###### #####
 #     #  #  #    #      #       #    # ##  ## #    # #    #   #   #      #    #
 #     # #    #   #      #       #    # # ## # #    # #    #   #   #####  #    #
 #   # # ######   #      #       #    # #    # #####  #    #   #   #      #####
 #    #  #    #   #      #     # #    # #    # #      #    #   #   #      #   #
  #### # #    #   #       #####   ####  #    # #       ####    #   ###### #    #

by Mickahell.
`

func DrawStart() {
	logger.GetLogger().LogDraw(draw)
	logger.GetLogger().LogDraw("Version : " + version_file + "\n")
}

func GetVersion() string {
	return version_file
}

func GetAppName() string {
	return app_name
}

func GetUsageMan() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()

	switch os.Args[0] {
	case ComputeStr:
		// do nothing aka pass
	default:
		fmt.Printf("Sub-commands available :\n %s \n", ComputeStr)
	}
}

func GetConfigPathMan() string {
	man := "Path for the configuration yaml file.\n" +
		"The configuration file can contains any options presented."

	return man
}

func GetLogLevelMan() string {
	man := "level of log to print :\n" +
		"0 : Informational (default value)\n" +
		"1 : Warning\n" +
		"2 : Error --> Always shown\n" +
		"3 : Critical --> Always shown"

	return man
}

func GetComputePathMan() string {
	man := "Path of the project to compute."

	return man
}

func GetFileExeMan() string {
	man := "Name of the file to compute the project."

	return man
}

func GetRequirementsFileMan() string {
	man := "Name of the requirements file to setup the project."

	return man
}

func GetPythonVerMan() string {
	man := "Version of python to use."

	return man
}

func GetDebianPkgMan() string {
	man := "List of debian package to install."

	return man
}
