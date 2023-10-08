package on_go

import (
	"qat-computer/helpers"
	"qat-computer/logger"
	"strings"
)

func ExecSummary() {
	// Debian packages setup
	sourcelist_update := "apt update -yq"
	logger.GetLogger().LogInfo("on_go", BashCMD(sourcelist_update))
	if len(helpers.TheAppConfig().DebianPkg) > 0 {
		logger.GetLogger().LogDraw("####################\n" + "# Setup Debian deps...")
		debian_req := "apt install -yq " + strings.Join(helpers.TheAppConfig().DebianPkg, " ")
		logger.GetLogger().LogInfo("on_go", BashCMD(debian_req))
	}

	// Python package setup
	logger.GetLogger().LogDraw("####################\n" + "# Setup Python env...")
	python_setup := "apt install -yq " + helpers.TheAppConfig().PythonVer
	logger.GetLogger().LogInfo("on_go", BashCMD(python_setup))

	// Python program compute
	if len(helpers.TheAppConfig().ComputePath) > 0 {
		logger.GetLogger().LogDraw("####################\n" + "# Exec Python project...")
		python_exec := helpers.TheAppConfig().PythonVer + " " + helpers.TheAppConfig().ComputePath
		logger.GetLogger().LogInfo("on_go", BashCMD(python_exec))
	}
}
