package main

import (
	"os"
)

type Config struct {
	Period    string
	BrickURL  string
	Namespace string
	LogLevel  string
}

var config Config

// configmap -> env 읽어서 객체로 저장
func ConfigSetting() {

	config.Period = os.Getenv("period")
	config.LogLevel = os.Getenv("log_level")
	config.BrickURL = "/api/container/info"
	config.Namespace = os.Getenv("namespace")

}

func GetConfig() Config {
	ConfigSetting()
	return config
}
