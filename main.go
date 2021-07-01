package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyobins/service-discovery/cmd"
	"github.com/sirupsen/logrus"
)

func http

func main() {
	config := GetConig()
	r := mux.NewRouter()
	r.HandleFunc("/",
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
