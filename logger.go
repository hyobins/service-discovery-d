package main

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// log객체 setting
func init() {
	logger = logrus.New()
	// logger.SetLevel(config.LogLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func GetLogger() *logrus.Logger {
	return logger
}
