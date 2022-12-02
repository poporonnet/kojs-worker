package cmd

import (
	"flag"
	"github.com/mct-joken/jkojs-worker/pkg/util"
)

func Command() {
	var langFlag = flag.String("lang", "unset", "処理系名")
	var ID = flag.String("id", "0", "問題ID")
	var pretty = flag.Bool("p", false, "true-> 整形表示")
	flag.Parse()

	util.InitLogger(*pretty)

	util.Logger.Debug("-----JK-OJS WORKER-----")
	util.Logger.Debug("Boot done")
	util.Logger.Debug("Copyright (C) 2022 National Institute of Technology, Matsue College Information Club (松江高専 情報科学研究部)\n")

	util.Logger.Sugar().Debugf("Config load done: language type: %s, problem ID: %s\n", *langFlag, *ID)

	decideLanguage(*langFlag, *ID)
}
