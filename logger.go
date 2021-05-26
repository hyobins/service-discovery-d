package main

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter()
}

func GetLogger() *logrus.Logger {
	return logger
}

//for testing
//func SetLogger(1 *logrus.Logger) {
//	logger = 1
//}
