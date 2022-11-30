package cmd

import (
	"fmt"
	"github.com/mct-joken/jkojs-worker/pkg/compiler"
	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func decideLanguage(langType string, problemID string) {
	v, found := util.LANGUAGE[langType]
	if !found {
		return
	}
	fmt.Println(v)

	err := compiler.Compile(v)
	if err != nil {
		return
	}
}
