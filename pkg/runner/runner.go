package runner

import (
	"context"
	"fmt"
	"github.com/mct-joken/jkojs-worker/pkg/util"
	"os/exec"
)

func Run(l util.Language) error {
	cmd := fmt.Sprintf(l.ExecCMD, "./built", "./case", "test.txt")
	util.Logger.Sugar().Debugf("Exec Command: %s\n", cmd)
	res, err := exec.CommandContext(context.Background(), "sh", "-c", cmd).CombinedOutput()
	util.Logger.Sugar().Debugf("Exec Result: %s, Exec Error: %v\n", string(res), err)

	if err != nil {
		return err
	}
	return nil
}
