package main

import (
	"os"

	"github.com/hyobins/service-discovery/cmd"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	cmd.Execute()
	log.SetOutput(os.Stdout)
	Formatter := new(logrus.TextFormatter)
	Formatter.DisableColors = true
	Formatter.FullTimestamp = true
	Formatter.PadLevelText = true
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}
