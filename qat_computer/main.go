package main

import (
	"fmt"
	"os"
	"runtime"
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

	if flags.QiskitVersionflag {
		qiskit_vers := []string{
			"python3",
			"-c",
			"from qiskit import __qiskit_version__; print(__qiskit_version__)",
		}
		logger.GetLogger().LogDraw(on_go.RunCMD(qiskit_vers).Stdout)
		os.Exit(0)
	}

	if flags.OSVersionflag {
		run_os := runtime.GOOS
		switch run_os {
		case "darwin":
			logger.GetLogger().LogDraw(on_go.RunCMD([]string{"sw_vers"}).Stdout)
		case "linux":
			logger.GetLogger().LogDraw(on_go.RunCMD([]string{"lsb_release", "-cdr"}).Stdout)
		default:
			fmt.Printf("%s.\n", run_os)
		}
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
		on_go.ExecSummary()
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
	logger.GetLogger().LogDraw("####################\n" + "# End.")
}
