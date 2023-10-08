package main

import (
	"flag"
	"os"
	"sync"

	"qat-computer/logger"
	"qat-computer/on_go"
	"qat-computer/utils"

	"qat-computer/docs"
	"qat-computer/helpers"
)

var Versionflag bool
var Configflag bool

func startoptions() {
	flag.BoolVar(
		&Versionflag,
		"version", false,
		"print version.",
	)

	flag.BoolVar(
		&Configflag,
		"show-config", false,
		"print config.",
	)

	flag.IntVar(
		&helpers.LogLevelflag,
		"log-level", 0,
		docs.GetLogLevelMan(),
	)

	flag.StringVar(
		&helpers.ComputePathflag,
		"compute", "",
		docs.GetComputePathMan(),
	)

	flag.StringVar(
		&helpers.PythonVerflag,
		"python-version", "python3.10",
		docs.GetPythonVerMan(),
	)

	flag.StringVar(
		&helpers.DebianPkgflag,
		"debian-pkg", "",
		docs.GetDebianPkgMan(),
	)

	flag.StringVar(
		&helpers.ConfPathflag,
		"conf", "",
		docs.GetConfigPathMan(),
	)

	flag.Parse()
}

func InitConf() {
	helpers.InitFile()
	logger.GetLogger().LogInit()
}

func main() {
	startoptions()
	InitConf()

	if Versionflag {
		logger.GetLogger().LogDraw(docs.GetVersion())
		os.Exit(0)
	}

	if Configflag {
		logger.GetLogger().LogDraw(utils.ToJSON(helpers.TheAppConfig()))
		os.Exit(0)
	}

	// pretty print
	docs.DrawStart()
	logger.GetLogger().LogInfo("main", "Qat ready to compute.")

	// create a WaitGroup
	wg := new(sync.WaitGroup)
	// add goroutines to `wg` WaitGroup
	wg.Add(1)

	go func() {
		// run func
		// logger.GetLogger().
		//	LogCritical("main", "listen cas error", cocas.GetCasServer().ListenAndServe())
		on_go.ExecSummary()
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
	logger.GetLogger().LogDraw("####################\n" + "# End.")
}
