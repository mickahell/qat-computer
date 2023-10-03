package main

import (
	"flag"
	"os"
	"sync"

	"cat-computer/logger"
	"cat-computer/utils"

	"cat-computer/docs"
	"cat-computer/helpers"
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

	flag.StringVar(
		&helpers.Confpathflag,
		"conf", "",
		docs.GetConfigPathMan(),
	)

	flag.IntVar(
		&helpers.LogLevelflag,
		"log-level", 0,
		docs.GetLogLevelMan(),
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
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
