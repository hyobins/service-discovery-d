package log

import (
	"fmt"
	"os"
	"path"
	"runtime"

	logrus "github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger() *logrus.Logger {
	Logger := logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetOutput(os.Stdout)

	Formatter := new(logrus.TextFormatter)

	Formatter.DisableColors = true
	Formatter.FullTimestamp = true
	Formatter.PadLevelText = true
	Formatter.FullTimestamp = true
	Formatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
		_, filename := path.Split(f.File)
		retFile := fmt.Sprintf("[ %s:%d ]", filename, f.Line)
		return "", retFile
	}
	Logger.SetFormatter(Formatter)
	return Logger
}

func GetLogger() *logrus.Logger {
	return Logger
}

func SetLogLevel(lv string) {
	switch lv {
	case "debug", "d":
		Logger.SetLevel(logrus.DebugLevel)
	case "info", "i":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn", "w":
		Logger.SetLevel(logrus.WarnLevel)
	case "err", "e":
		Logger.SetLevel(logrus.ErrorLevel)
	}
}

func GetLogLevel() string {
	switch Logger.GetLevel() {
	case logrus.DebugLevel:
		return "debug"
	case logrus.InfoLevel:
		return "info"
	case logrus.WarnLevel:
		return "warn"
	case logrus.ErrorLevel:
		return "error"
	}
	return "unknown"
}

type LogrusWriter struct {
	Logger logrus.FieldLogger
}

// log객체 setting
//func LogSetting() {
//	logger = logrus.New()
//	// logger.SetLevel(config.LogLevel)
//	logger.SetFormatter(&logrus.TextFormatter{
//		DisableColors: true,
//		FullTimestamp: true,
//	})
//}
//
//func GetLogger() *logrus.Logger {
//	LogSetting()
//	return logger
//}
