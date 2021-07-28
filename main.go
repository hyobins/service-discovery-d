package main

import (
	"github.com/hyobins/service-discovery/cmd"
)

// Logger iris-cloud logger
//var Logger *logrus.Logger

//func LogInit() {
//	Logger = logrus.New()
//	Logger.SetReportCaller(true)
//	Logger.SetOutput(os.Stdout)
//	Logger.SetLevel(logrus.DebugLevel)

//	Formatter := new(logrus.JBFormatter)

/* Set Print Timestamp And Format */
// Formatter.DisableColors = true
// Formatter.FullTimestamp = true
// Formatter.PadLevelText = true
// Formatter.DisableLevelTruncation = true // bool
//	Formatter.TimestampFormat = "2006-01-02 15:04:05.000"
// Formatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
//  _, filename := path.Split(f.File)
//  retFile := fmt.Sprintf("[ %s:%d ]", filename, f.Line)
//  return "", retFile
// }
//	Logger.SetFormatter(Formatter)
//}

func main() {
	//LogInit()
	//Logger.Infof("print test")
	cmd.Execute()
}
