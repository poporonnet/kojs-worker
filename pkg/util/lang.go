package util

/*
	使用可能な言語リスト

	- C(GCC/Clang)
	- C++ (G++/Clang++)
	- Ruby3
	- JavaScript (Node.js)
	- Golang (Golang-Go)
	- Python3

	- [TBD] Crystal
	- [TBD] C# (mono)
	- [TBD] TypeScript (Node.js-TSC)
	- [TBD] Java (JVM)
	- [TBD] Rust (Rustc)
*/

type Language struct {
	Name       string // 言語名 e.g: C++
	Type       string // コンパイラタイプ e.g: GCC
	CompileCMD string // コンパイルコマンド e.g: gcc main.c
	ExecCMD    string // 実行コマンド e.g: ./a.out
}

// 言語リスト
var LANGUAGE = map[string]Language{
	"GCC": {
		Name:       "C",
		Type:       "GCC",
		CompileCMD: "gcc -std=gnu11 -Wall -Wextra -O2 %s/main.c -o %s/a.out",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	"Clang": {
		Name:       "C",
		Type:       "Clang",
		CompileCMD: "clang -std=c11 -Wall -Wextra -O2 %s/main.c -o %s/a.out",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	"G++": {
		Name:       "C++",
		Type:       "G++",
		CompileCMD: "g++ -std=gnu++17 -Wall -Wextra -O2 %s/main.cpp -o %s/a.out",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	"Clang++": {
		Name:       "C++",
		Type:       "Clang++",
		CompileCMD: "clang++ -std=c++17 -Wall -O2 %s/main.cpp -o %s/a.out",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	"Ruby": {
		Name:       "Ruby",
		Type:       "Ruby",
		CompileCMD: "ruby -w -c %s/main.rb; echo %s > /dev/null",
		ExecCMD:    "ruby %s/main.rb < %s/%s",
	},
}
