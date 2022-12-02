package util

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(dev bool) *zap.Logger {
	if dev {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}
	return Logger
}
