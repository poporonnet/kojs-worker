package cmd

import (
	"flag"
	"fmt"
)

func Command() {
	var langFlag = flag.String("lang", "GCC", "処理系名")
	var ID = flag.String("id", "0", "問題ID")
	flag.Parse()

	fmt.Println(*langFlag, *ID)
	decideLanguage(*langFlag, *ID)
}
