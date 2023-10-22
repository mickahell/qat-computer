package docs

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed VERSION.txt
var version_file string

const app_name string = "Qat Command"
const BinaryName string = "qat-computer"
const ImageName string = "qatcomputer"
const RegistryImage string = "ghcr.io/mickahell/" + ImageName + "-dev"
const RegistryTag string = "latest"
const ContainerName string = "qatcomputer"

const PullStr = "pull"
const InitStr = "init"
const ComputerStr = "qatcomputer"
const StopStr = "stop"

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
	case InitStr:
		// do nothing aka pass
	case ComputerStr:
		// do nothing aka pass
	default:
		fmt.Printf(
			"Sub-commands available :\n %s, %s, %s, %s \n",
			PullStr,
			InitStr,
			StopStr,
			ComputerStr,
		)
	}
}
