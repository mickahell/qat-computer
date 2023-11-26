package flags

import (
	"flag"
	"fmt"
	"os"
	"qat-cmd/docs"
)

var Initflag bool
var Computerflag bool
var Pullflag bool
var Stopflag bool
var StopRemoveflag bool
var ComputeVolumeflag string
var ConfVolumeflag string
var ComputerArgsflag string
var FullImageNameflag string
var ImageTagflag string

var pullCMD = flag.NewFlagSet(docs.PullStr, flag.ExitOnError)
var initCMD = flag.NewFlagSet(docs.InitStr, flag.ExitOnError)
var computerCMD = flag.NewFlagSet(docs.ComputerStr, flag.ExitOnError)
var stopCMD = flag.NewFlagSet(docs.StopStr, flag.ExitOnError)

func setupPullFlags() {
	for _, fs := range []*flag.FlagSet{pullCMD} {
		fs.StringVar(
			&FullImageNameflag,
			"image", docs.RegistryImage,
			"full url of the image to pull.",
		)

		fs.StringVar(
			&ImageTagflag,
			"tag", docs.RegistryTag,
			"image tag to pull.",
		)
	}
}

func setupInitFlags() {
	for _, fs := range []*flag.FlagSet{initCMD} {
		fs.StringVar(
			&FullImageNameflag,
			"image", docs.RegistryImage,
			"full url of the image to pull.",
		)

		fs.StringVar(
			&ComputeVolumeflag,
			"compute-volume", "",
			"compute path for the compute volume.",
		)

		fs.StringVar(
			&ConfVolumeflag,
			"conf-volume", "",
			"conf path for the conf volume.",
		)
	}
}

func setupComputerFlags() {
	for _, fs := range []*flag.FlagSet{computerCMD} {
		fs.StringVar(
			&ComputerArgsflag,
			"computer-args", "--help",
			"any args for qatcomputer binary.",
		)
	}
}

func setupStopFlags() {
	for _, fs := range []*flag.FlagSet{stopCMD} {
		fs.BoolVar(
			&StopRemoveflag,
			"remove", false,
			"remove the qatcomputer container.",
		)
	}
}

func StartOptions() {
	setupPullFlags()
	setupInitFlags()
	setupComputerFlags()
	setupStopFlags()

	flag.Usage = docs.GetUsageMan

	switch os.Args[1] {
	case docs.PullStr:
		Pullflag = true
		err := pullCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	case docs.InitStr:
		Initflag = true
		err := initCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	case docs.ComputerStr:
		Computerflag = true
		err := computerCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	case docs.StopStr:
		Stopflag = true
		err := stopCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	default:
		flag.Parse()
	}
}
