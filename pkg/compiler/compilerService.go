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
	fmt.Println(string(res), err)
	return nil
}
