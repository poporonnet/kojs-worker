package runner

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mct-joken/jkojs-worker/pkg/util"
)

// ProblemConfig 問題ごとの設定
type ProblemConfig struct {
	ID          string   `json:"id"`
	TimeLimit   int      `json:"timeLimit"`   // 実行時間制限
	MemoryLimit int      `json:"memoryLimit"` // メモリ制限
	CaseFiles   []string `json:"caseFiles"`   // ケースファイルのファイルパス
}

const CONFIGPATH = "./case"

func LoadConfig(id string) (ProblemConfig, error) {
	path := fmt.Sprintf("%s/%s.json", CONFIGPATH, id)
	f, err := os.ReadFile(path)
	if err != nil {
		util.Logger.Sugar().Fatalf("Failed to read Config file: %s", err.Error())
		return ProblemConfig{}, err
	}

	config := ProblemConfig{}
	err = json.Unmarshal(f, &config)
	if err != nil {
		return ProblemConfig{}, err
	}
	util.Logger.Sugar().Debugf("%v", config)

	return config, nil
}
