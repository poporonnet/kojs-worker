package runner

import (
	"context"
	"fmt"
	"github.com/mct-joken/jkojs-worker/pkg/util"
	"os/exec"
)

func Run(l util.Language) error {
	cmd := fmt.Sprintf(l.ExecCMD, "./built", "./case", "test.txt")
	fmt.Println(cmd)
	res, err := exec.CommandContext(context.Background(), "sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}
