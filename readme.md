# JK-OJS (Worker)
JK-OJS の Worker (コードをコンパイル/実行し,解となる入力を与え結果を返す)部分です.

## Requirements
- Go 1.19.x
- Linux x86_64
> **Warning**
> Linux向けビルドでないとコンパイルできません.

## 使い方
### 準備
まず,これらのファイルを配置する必要があります
- `case/` ケースファイル,問題の制約ファイル( `<problemID>.json` )
- `test/` 実行するプログラムファイル( `Ruby` 以外)
- `built/` `Ruby` のプログラムファイルは事前にこのディレクトリに配置する必要があります.

### 実行
このように与えると動作します.
```
go build -o jkworker main.go
./jkworker -lang <langID> -id <problemID>
```
最後に `-p` をつけるとデバッグログを吐きます.  

## 対応言語表
凡例: `<言語名> (<言語ID>)`
- C (GCC/Clang)
- C++ (G++/Clang++)
- Ruby (Ruby)

