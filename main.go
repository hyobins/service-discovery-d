package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hee")
	config := GetConfig()

	//로그 객체 초기화
	logger := GetLogger()

	logger.WithFields(logrus.Fields{
		"period":    config.Period,
		"namespace": config.Namespace,
		"brickURL":  config.BrickURL,
	}).Info("Configuration Information")

	fmt.Println(os.Getenv("env"))

}
