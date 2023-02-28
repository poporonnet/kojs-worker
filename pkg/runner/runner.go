package runner

import (
	"context"
	"strings"
	"syscall"

	"time"

	"github.com/k1LoW/exec"

	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func Run(l util.Language, status *util.ExecuteStatus, cfg ProblemConfig) error {

	for _, v := range cfg.CaseFiles {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TimeLimit)*time.Millisecond)
		defer cancel()

		command := l.ExecCMD("./built", "./case", v)
		// 実行(実行時間制限付き): configにあるミリ秒経過したら実行中でもkillする
		cmd := exec.CommandContext(ctx, "sh", "-c", command)

		// コマンドを実行する stdout/stderrが返る
		res, err := cmd.Output()

		// 実行時間(ミリ秒)
		duration := cmd.ProcessState.UserTime().Milliseconds()
		// syscall.Rusageは*nixにしか無いメソッド Windowsでは実行できない
		memUsage := (cmd.ProcessState.SysUsage().(*syscall.Rusage)).Maxrss / 4
		exitStatus := cmd.ProcessState.ExitCode()

		util.Logger.Sugar().Debugf("Exec Command: %s\n", command)
		util.Logger.Sugar().Debugf("Exec Result: %s, Exec Error: %v\n", string(res), err)

		status.Results = append(status.Results, util.CaseResult{
			CaseID:      strings.Split(v, ".")[0],
			Output:      string(res),
			ExitStatus:  exitStatus, // 終了コードが-1のときはTLEで強制終了されている
			Duration:    int(duration),
			MemoryUsage: int(memUsage),
		})

		if err != nil {
			return err
		}

	}

	return nil
}
