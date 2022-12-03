package cmd

import (
	"encoding/json"
	"os"

	"github.com/mct-joken/jkojs-worker/pkg/compiler"
	"github.com/mct-joken/jkojs-worker/pkg/runner"
	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func decideLanguage(langType string, problemID string) {
	// 言語名から実行すべきコマンドを探す
	v, found := util.LANGUAGE[langType]
	if !found {
		return
	}

	status := util.ExecuteStatus{
		SubmissionID: "",
		ProblemID:    problemID,
		LanguageType: langType,
	}

	err := compiler.Compile(v, &status)
	if err != nil {
		return
	}

	cfg, err := runner.LoadConfig(problemID)
	if err != nil {
		return
	}

	runner.Run(v, &status, cfg)

	resultOutPut, _ := json.Marshal(status)
	err = os.WriteFile("out.json", resultOutPut, 0666)
	if err != nil {
		util.Logger.Sugar().Fatalf("ファイルの出力に失敗: %v", err.Error())
		return
	}

}
