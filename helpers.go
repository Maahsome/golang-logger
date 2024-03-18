package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Used to format the Prettyfier in logrus.TextFormatter. This will format the Reporter call on the file.
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return fmt.Sprintf(" | %v", arr[len(arr)-1])
}

// Used to format the Prettyfier in logrus.TextFormatter. This will format the Reporter call on the function.
func formatFunction(path string) string {
	arr := strings.Split(path, "/")
	return fmt.Sprintf(" | %v |", arr[len(arr)-1])
}

// Used to format the Prettyfier in logrus.TextFormatter. This will format the Reporter call on the stdout.
func formatStdOut(path string) string {
	arr := strings.Split(path, "/")
	return fmt.Sprintf("%v", arr[len(arr)-1])
}

// Configure a logger with some preset formatting options.
func formatter(logger *logrus.Logger) *logrus.Logger {
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "15:04:05",
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return formatFunction(f.Function), fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)
	return logger
}
