package main

import (
	"os"
)

type Config struct {
	Period    string `yaml: "period"`
	BrickURL  string `yaml: "brick_url"`
	Namespace string `yaml: "namespace"`
	LogLevel  string `yaml: "loglevel"`
}

var config Config

// configmap -> env 읽어서 객체로 저장
func main() {

	config.Period = os.Getenv("period")
	config.LogLevel = os.Getenv("log_level")
	config.BrickURL = "/api/container/info"
	config.Namespace = os.Getenv("namespace")

}

func GetConfig() Config {
	return config
}
