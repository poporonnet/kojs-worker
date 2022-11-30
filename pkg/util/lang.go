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

var LANGUAGE = []Language{
	{
		Name:       "C",
		Type:       "GCC",
		CompileCMD: "gcc -std=gnu11 -O2 -o %s/a.out %s/%s",
		ExecCMD:    "%s/a.out",
	},
	{
		Name:       "C",
		Type:       "Clang",
		CompileCMD: "clang -std=c11 -O2 -o %s/a.out %s/%s",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	{
		Name:       "C++",
		Type:       "G++",
		CompileCMD: "g++ -std=gnu++17 -Wall -Wextra -O2 -o %s/a.out %s/%s",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	{
		Name:       "C++",
		Type:       "Clang++",
		CompileCMD: "clang++ -std=c++17 -stdlib=libc++ -Wall -O2 -o %s/a.out %s/%s",
		ExecCMD:    "%s/a.out < %s/%s",
	},
	{
		Name:       "Ruby",
		Type:       "Ruby",
		CompileCMD: "ruby -w -c %s/%s",
		ExecCMD:    "ruby %s/%s < %s/%s",
	},
}
