package logger

import (
	"fmt"
	"github.com/RianIhsan/go-clean-architecture-v2/config"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

var fieldsLogrusLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

func NewLogrusLogger(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	level := getLogrusLevel(cfg)

	var formatter logrus.Formatter
	if cfg.Logger.Encoding == "text" {
		formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
			DisableColors:   false,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				file = fmt.Sprintf("%s:%d", frame.File, frame.Line)
				return function, filepath.Base(file)
			},
		}
	} else {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
			PrettyPrint:     true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				file = fmt.Sprintf("%s:%d", frame.File, frame.Line)
				return function, filepath.Base(file)
			},
		}
	}

	logger.SetLevel(level)
	logger.SetFormatter(formatter)
	logger.SetReportCaller(cfg.Logger.Caller)
	logger.SetOutput(os.Stdout)

	return logger
}

func getLogrusLevel(cfg *config.Config) logrus.Level {
	level, exist := fieldsLogrusLevelMap[cfg.Logger.Level]
	if !exist {
		level = logrus.DebugLevel
	}
	return level
}
