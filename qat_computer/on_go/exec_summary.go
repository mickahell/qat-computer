package on_go

import (
	"path/filepath"
	"qat-computer/helpers"
	"qat-computer/logger"
	"qat-computer/utils"
	"strings"
)

func ExecSummary() {
	// Debian & Python setup
	if len(helpers.TheAppConfig().DebianPkg) != 0 &&
		BashCMD(helpers.TheAppConfig().PythonVer+" --version").Stderr != nil {
		logger.GetLogger().LogDraw("####################\n" + "# Setup Python & Debian deps...")
		sourcelist_update := "apt update -yq"
		logger.GetLogger().LogInfo("on_go", BashCMD(sourcelist_update).Stdout)
		debian_req := "apt install -yq " + strings.Join(helpers.TheAppConfig().DebianPkg, " ")
		logger.GetLogger().LogInfo("on_go", BashCMD(debian_req).Stdout)
		python_setup := "apt install -yq " + helpers.TheAppConfig().PythonVer
		logger.GetLogger().LogInfo("on_go", BashCMD(python_setup).Stdout)
	}

	// Python package setup
	logger.GetLogger().LogDraw("####################\n" + "# Setup Python env...")
	if utils.FileExists(filepath.Join(helpers.TheAppConfig().ComputePath, "setup.py")) ||
		utils.FileExists(filepath.Join(helpers.TheAppConfig().ComputePath, "pyproject.toml")) {
		project_setup := "pip install " + filepath.Join(helpers.TheAppConfig().ComputePath, ".")
		logger.GetLogger().LogInfo("on_go", BashCMD(project_setup).Stdout)
	}
	if helpers.TheAppConfig().RequirementsFile != "" {
		requirements_setup := "pip install -r " + filepath.Join(
			helpers.TheAppConfig().ComputePath,
			helpers.TheAppConfig().RequirementsFile,
		)
		logger.GetLogger().LogInfo("on_go", BashCMD(requirements_setup).Stdout)
	}

	// Python program compute
	if len(helpers.TheAppConfig().ComputePath) > 0 {
		logger.GetLogger().LogDraw("####################\n" + "# Exec Python project...")
		python_exec := helpers.TheAppConfig().PythonVer + " " + filepath.Join(
			helpers.TheAppConfig().ComputePath,
			helpers.TheAppConfig().FileExeName,
		)
		logger.GetLogger().LogInfo("on_go", BashCMD(python_exec).Stdout)
	}
}
