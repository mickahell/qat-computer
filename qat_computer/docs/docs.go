package docs

import (
	_ "embed"

	"qat-computer/logger"
)

//go:embed VERSION.txt
var version_file string
var app_name string = "Qat Computer"

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

func GetConfigPathMan() string {
	man := "Relative path for the configuration yaml file.\n" +
		"The configuration file can contains any options presented below."

	return man
}

func GetLogLevelMan() string {
	man := "level of log to print :\n" +
		"0 : Informational\n" +
		"1 : Warning\n" +
		"2 : Error --> Always shown\n" +
		"3 : Critical --> Always shown"

	return man
}

func GetComputePathMan() string {
	man := "Relative path for the project to compute yaml file"

	return man
}
