package model

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Period    string
	BrickURL  string
	Namespace string
	LogLevel  logrus.Level
}

var config Config

// configmap -> env 읽어서 객체로 저장
func ConfigSetting() {
	config.Period = os.Getenv("PERIOD")

	// log_level := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	log_level := "WARN"
	if log_level == "INFO" {
		config.LogLevel = logrus.InfoLevel
	} else if log_level == "WARN" {
		config.LogLevel = logrus.WarnLevel
	}
	config.BrickURL = "/api/container/info"
	//config.Namespace = os.Getenv("NAMESPACE")
	config.Namespace = "iris-cloud"
}

func GetConfig() Config {
	ConfigSetting()
	return config
}
