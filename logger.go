package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Create a logger object with a set level of INFO, reporter toggle, and customized textformatter.
func CreateStandardLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger = formatter(logger)
	return logger
}

// Create a error logger object with a set level of ERROR, reporter toggle, customized textformatter, and a path to a log file.
func CreateErrorLogger(logFilePath string) (*logrus.Logger, error) {
	// Setup error logger
	_, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &logrus.Logger{}, errors.Wrap(err, "unable to open file")
	}
	errorLog := logrus.New()
	// Set custom formatting
	formatter(errorLog)
	// Allow the logger to write ErrorLevel to an log file
	pathMap := lfshook.PathMap{
		logrus.ErrorLevel: logFilePath,
	}
	errorLog.SetReportCaller(true)
	// Set a hook to format the content going into the log file
	errorLog.Hooks.Add(lfshook.NewHook(pathMap, &logrus.TextFormatter{
		TimestampFormat:        "15:04:05",
		ForceColors:            false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return formatStdOut(f.Function), fmt.Sprintf("%s:%d", formatStdOut(f.File), f.Line)
		},
	}))
	errorLog.SetLevel(logrus.ErrorLevel)

	return errorLog, nil
}

// Parses string to logrus level object
// Default log level is logrus.DebugLevel
func LogLevelFromString(logLevelString string) logrus.Level {
	ll := logrus.InfoLevel
	switch strings.ToLower(logLevelString) {
	case "trace":
		ll = logrus.TraceLevel
	case "debug":
		ll = logrus.DebugLevel
	case "info":
		ll = logrus.InfoLevel
	case "warning":
		ll = logrus.WarnLevel
	case "error":
		ll = logrus.ErrorLevel
	case "fatal":
		ll = logrus.FatalLevel
	default:
		ll = logrus.DebugLevel
	}
	return ll
}
