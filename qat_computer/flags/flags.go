package flags

import (
	"flag"
	"fmt"
	"os"
	"qat-computer/docs"
	"qat-computer/helpers"
)

var Versionflag bool
var QiskitVersionflag bool
var OSVersionflag bool
var Configflag bool

var computeCmd = flag.NewFlagSet(docs.ComputeStr, flag.ExitOnError)

func setupCommonFlags() {
	for _, fs := range []*flag.FlagSet{computeCmd} {
		fs.IntVar(
			&helpers.LogLevelflag,
			"log-level", 0,
			docs.GetLogLevelMan(),
		)

		fs.StringVar(
			&helpers.RequirementsFileflag,
			"requirements", "",
			docs.GetRequirementsFileMan(),
		)

		fs.StringVar(
			&helpers.PythonVerflag,
			"python-version", "python3",
			docs.GetPythonVerMan(),
		)

		fs.StringVar(
			&helpers.DebianPkgflag,
			"debian-pkg", "",
			docs.GetDebianPkgMan(),
		)

		fs.StringVar(
			&helpers.ConfPathflag,
			"conf", "",
			docs.GetConfigPathMan(),
		)
	}
}

func setupComputeFlags() {
	for _, fs := range []*flag.FlagSet{computeCmd} {
		fs.StringVar(
			&helpers.ComputePathflag,
			"compute", "",
			docs.GetComputePathMan(),
		)

		fs.StringVar(
			&helpers.FileExeNameflag,
			"file_exe", "main.py",
			docs.GetFileExeMan(),
		)
	}
}

func StartOptions() {
	setupCommonFlags()
	setupComputeFlags()

	flag.BoolVar(
		&Versionflag,
		"version", false,
		"print version.",
	)

	flag.BoolVar(
		&QiskitVersionflag,
		"qiskit-version", false,
		"print qiskit version.",
	)

	flag.BoolVar(
		&OSVersionflag,
		"os-version", false,
		"print os version.",
	)

	flag.BoolVar(
		&Configflag,
		"show-config", false,
		"print config.",
	)

	flag.StringVar(
		&helpers.ConfPathflag,
		"conf", "",
		docs.GetConfigPathMan(),
	)

	flag.Usage = docs.GetUsageMan

	switch os.Args[1] {
	case docs.ComputeStr:
		err := computeCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	default:
		flag.Parse()
	}
}
