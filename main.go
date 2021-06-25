package main

import (
	"github.com/hyobins/service-discovery/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	config := GetConfig()

	//로그 객체 초기화
	logger := GetLogger()
	// logger.SetLevel(config.LogLevel)
	logger.WithFields(logrus.Fields{
		"period":    config.Period,
		"namespace": config.Namespace,
		"brickURL":  config.BrickURL,
	}).Info("Configuration Information")

	cmd.Execute()
}
