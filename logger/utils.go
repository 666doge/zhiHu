package logger

import (
	"runtime"
	"time"
	"fmt"
)

type LogInfo struct {
	IsWarn bool
	LogMsg string
}

func GetLineInfo() (fileName string, fnName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		lineNo = line
		fnName = runtime.FuncForPC(pc).Name()
	}
	return
}

func  GetlogInfo(level int, format string, args ...interface{}) *LogInfo {
	nowString := time.Now().Format("2006-01-02 15:04:05")
	fileName, funcName, lineNo := GetLineInfo()
	msg := fmt.Sprintf(format, args...)

	logMsg := fmt.Sprintf(
		"%s [%s] %s %s:%d %s\n",
		nowString,
		LogLevelText(level),
		fileName,
		funcName,
		lineNo,
		msg,
	)

	return &LogInfo{
		IsWarn: level >= LogLevelWarn,
		LogMsg: logMsg,
	}
}