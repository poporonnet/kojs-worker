package compiler

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func Compile(lang util.Language, status *util.ExecuteStatus) error {
	cmd := fmt.Sprintf(lang.CompileCMD, "./test", "./built")
	res, err := exec.CommandContext(context.Background(), "sh", "-c", cmd).CombinedOutput()

	util.Logger.Sugar().Debugf("Compiler Command: %s\n", cmd)
	util.Logger.Sugar().Debugf("Compiler Message: %s, Command Error: %v", string(res), err)

	status.CompilerMessage = string(res)
	if err != nil {
		status.CompileErrorMessage = err.Error()
	}

	return err
}
