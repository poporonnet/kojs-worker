package cmd

import (
	"github.com/mct-joken/jkojs-worker/pkg/compiler"
	"github.com/mct-joken/jkojs-worker/pkg/runner"
	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func decideLanguage(langType string, problemID string, pretty bool) {
	v, found := util.LANGUAGE[langType]
	if !found {
		return
	}

	err := compiler.Compile(v)
	if err != nil {
		util.Logger.Sugar().Fatalf("Fatal Error: %s", err)
		return
	}
	err = runner.Run(v)
	if err != nil {
		util.Logger.Sugar().Fatalf("Fatal Error: %s", err)
		return
	}
}
