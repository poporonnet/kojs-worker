package compiler

import (
	"context"
	"fmt"
	"github.com/mct-joken/jkojs-worker/pkg/util"
	"os/exec"
)

func Compile(lang util.Language) error {
	cmd := fmt.Sprintf(lang.CompileCMD, "./test", "./built")
	res, err := exec.CommandContext(context.Background(), "sh", "-c", cmd).CombinedOutput()
	util.Logger.Sugar().Debugf("Compiler Command: %s\n", cmd)
	util.Logger.Sugar().Debugf("Compiler Message: %s, Compiler Error: %v", string(res), err)

	return nil
}
