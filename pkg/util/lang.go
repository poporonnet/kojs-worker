package util

import (
	"fmt"
)

/*
	使用可能な言語リスト

	- C(GCC/Clang)
	- C++ (G++/Clang++)
	- Ruby3
	- Golang (Golang-Go)
	- [ToDo] Python3
	- [ToDo] JavaScript (Node.js)
	- [ToDo] Java (JVM)

	- [TBD] Crystal
	- [TBD] C# (mono)
	- [TBD] TypeScript (Node.js-TSC)
	- [TBD] Rust (Rustc)
*/

type Language struct {
	Name       string                                                        // 言語名 e.g: C++
	Type       string                                                        // コンパイラタイプ e.g: GCC
	CompileCMD func(i string, o string) string                               // コンパイルコマンド e.g: gcc main.c
	ExecCMD    func(exec string, caseDir string, caseFileName string) string // 実行コマンド e.g: ./a.out
}

// 言語リスト
var LANGUAGE = map[string]Language{
	"GCC": {
		Name: "C",
		Type: "GCC",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("gcc -std=gnu11 -Wall -Wextra -O2 %s/main.c -o %s/a.out", i, o)
		},
		ExecCMD: func(exec string, caseDir string, casePath string) string {
			return fmt.Sprintf("%s/a.out < %s/%s", exec, caseDir, casePath)
		},
	},
	"Clang": {
		Name: "C",
		Type: "Clang",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("clang -std=c11 -Wall -Wextra -O2 %s/main.c -o %s/a.out", i, o)
		},
		ExecCMD: func(exec string, caseDir string, casePath string) string {
			return fmt.Sprintf("%s/a.out < %s/%s", exec, caseDir, casePath)
		},
	},
	"G++": {
		Name: "C++",
		Type: "G++",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("g++ -std=gnu++17 -Wall -Wextra -O2 %s/main.cpp -o %s/a.out", i, o)
		},
		ExecCMD: func(exec string, caseDir string, casePath string) string {
			return fmt.Sprintf("%s/a.out < %s/%s", exec, caseDir, casePath)
		},
	},
	"Clang++": {
		Name: "C++",
		Type: "Clang++",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("clang++ -std=c++17 -Wall -O2 %s/main.cpp -o %s/a.out", i, o)
		},
		ExecCMD: func(exec string, caseDir string, casePath string) string {
			return fmt.Sprintf("%s/a.out < %s/%s", exec, caseDir, casePath)
		},
	},
	"Ruby": {
		Name: "Ruby",
		Type: "Ruby",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("ruby -w -c %s/main.rb; cp %s/main.rb %s/main.rb", i, i, o)
		},
		ExecCMD: func(exec string, caseDir string, casePath string) string {
			return fmt.Sprintf("ruby %s/main.rb < %s/%s", exec, caseDir, casePath)
		},
	},
	"Go": {
		Name: "Go",
		Type: "Go",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("CGO_ENABLED=0 go build -o %s/a.out %s/main.go", o, i)
		},
		ExecCMD: func(exec string, caseDir string, caseFileName string) string {
			return fmt.Sprintf("%s/a.out < %s/%s", exec, caseDir, caseFileName)
		},
	},
	"Python3": {
		Name: "Python3",
		Type: "Python3",
		CompileCMD: func(i string, o string) string {
			return fmt.Sprintf("cp %s/main.py %s/main.py", i, o)
		},
		ExecCMD: func(exec string, caseDir string, caseFileName string) string {
			return fmt.Sprintf("python3 %s/main.py < %s/%s", exec, caseDir, caseFileName)
		},
	},
}
