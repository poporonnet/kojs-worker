package util

// ExecuteStatus 提出ごとのステータス
type ExecuteStatus struct {
	SubmissionID string `json:"submissionID"` // 提出ID
	ProblemID    string `json:"problemID"`    // 問題ID
	LanguageType string `json:"languageType"` // 言語/処理系

	CompilerMessage     string `json:"compilerMessage"`     // コンパイラが出力した警告など
	CompileErrorMessage string `json:"compileErrorMessage"` // コンパイルエラー

	Results []CaseResult `json:"results"` // ケースごとのステータス
}

// CaseResult ケースごとのステータス
type CaseResult struct {
	CaseID      string `json:"caseID"`     // ケースID
	Output      string `json:"output"`     // プログラム出力
	ExitStatus  int    `json:"exitStatus"` // 終了コード
	Duration    int    `json:"duration"`   // 実行時間
	MemoryUsage int    `json:"usage"`      // メモリ使用量
}
