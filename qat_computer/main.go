package main

import (
	"os"
	"sync"

	"qat-computer/flags"
	"qat-computer/logger"
	"qat-computer/on_go"
	"qat-computer/utils"

	"qat-computer/docs"
	"qat-computer/helpers"
)

func InitConf() {
	helpers.InitFile()
	logger.GetLogger().LogInit()
}

func main() {
	flags.StartOptions()
	InitConf()

	if flags.Versionflag {
		logger.GetLogger().LogDraw(docs.GetVersion())
		os.Exit(0)
	}

	if flags.Configflag {
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
