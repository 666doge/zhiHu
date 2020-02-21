package logger

import (
	"os"
	"fmt"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger (config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("config not found : log_level")
		return
	}
	level := GetLogLevel(logLevel)
	logger = &ConsoleLogger{
		level: level,
	}
	return
}

func (c *ConsoleLogger) Init() {}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		c.level = LogLevelDebug
		return
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	logInfo := GetlogInfo(LogLevelDebug, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	logInfo := GetlogInfo(LogLevelTrace, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	logInfo := GetlogInfo(LogLevelInfo, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	logInfo := GetlogInfo(LogLevelWarn, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	logInfo := GetlogInfo(LogLevelError, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	logInfo := GetlogInfo(LogLevelFatal, format, args...)
	fmt.Fprint(os.Stdout, logInfo.LogMsg)
}

func (c *ConsoleLogger) Close() {
}