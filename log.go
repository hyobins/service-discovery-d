package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func LogSetting() *logrus.Logger {
	Logger := logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetOutput(os.Stdout)

	Formatter := new(logrus.TextFormatter)

	Formatter.DisableColors = true
	Formatter.FullTimestamp = true
	Formatter.PadLevelText = true
	Formatter.FullTimestamp = true
	Logger.SetFormatter(Formatter)

	return Logger
}

func GetLogger() *logrus.Logger {
	Logger := LogSetting()
	return Logger
}
