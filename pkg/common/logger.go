package common

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

var logLevel string = "debug"

type StructLogger struct {
	*logrus.Logger
}

func InitLogger() {

	//if err := os.MkdirAll("", 0755); err != nil {
	//	panic(fmt.Errorf("create log dir: %s error", ""))
	//}

	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return strconv.Itoa(frame.Line), frame.Function
		},
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyFile:  "filename",
			logrus.FieldKeyFunc:  "func",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
		},
	}

	// FIXME: 日志bug还没有解决

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(fmt.Errorf("parse log level error, format %s", logLevel))
	}

	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        false,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(level)

	Logger = &StructLogger{
		Logger: log,
	}

}
