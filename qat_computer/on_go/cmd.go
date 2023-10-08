package on_go

import (
	"errors"
	"os/exec"
	"qat-computer/logger"
	"strings"
)

type Cmds struct {
	Stdout string
	Stderr error
}

var Cmd Cmds

func BashCMD(cmd string) *Cmds {

	args := strings.Fields(cmd)
	to_run := exec.Command(args[0], args[1:]...)
	stdout, err := to_run.Output()

	if err != nil {
		error_send := errors.New(string(stdout) + "\n" + err.Error())
		logger.GetLogger().LogCritical("on_go", "error with cmd : "+cmd, error_send, false)
	}

	Cmd.Stdout = strings.TrimSuffix(string(stdout[:]), "\n")
	Cmd.Stderr = err

	return &Cmd
}
