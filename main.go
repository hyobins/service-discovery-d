package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyobins/service-discovery/cmd"
	"github.com/hyobins/service-discovery/internal/api"
	"github.com/sirupsen/logrus"
)

func main() {
	//Config Setting
	config := GetConfig()

	r := mux.NewRouter()
	api.Register(r, &api.Context{})
	http.ListenAndServe(":8080", r)

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
