package logger
import (
	"fmt"
)

var logger LogInterface

func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		logger, err = NewFileLogger(config)
	case "console":
		logger, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport log name: %s", name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	logger.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}