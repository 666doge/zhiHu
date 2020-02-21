package logger

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"strings"
)

type FileLogger struct {
	level int
	logPath string
	logName string
	file *os.File
	warnFile *os.File
	LogInfoChan chan *LogInfo
	logSplitType int
	logSplitSize int64
	lastSplitHour int
}

func NewFileLogger (config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("config not found : log_level")
		return
	}

	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("config not found : log_path")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("config not found : log_name")
		return
	}

	// 日志切分相关信息 - 按小时切分， 按文件大小切分
	logSplitTypeStr, ok := config["log_split_type"]
	var logSplitType int = LogSplitByHour
	var logSplitSize int64 = 104857600 // 默认大小为100M
	if ok {
		if logSplitTypeStr == "size" {
			logSplitType = LogSplitBySize
			logSplitSizeStr, ok := config["log_split_size"]
			if ok {
				size, err := strconv.ParseInt(logSplitSizeStr, 10, 64)
				if err == nil {
					logSplitSize = size
				}
			}
		} else if logSplitTypeStr == "hour" {
			logSplitType = LogSplitByHour
		}
	}

	level := GetLogLevel(logLevel)
	logger = &FileLogger{
		level: level,
		logPath: logPath,
		logName: logName,
		LogInfoChan: make(chan *LogInfo, 500),
		logSplitType: logSplitType,
		logSplitSize: logSplitSize,
		lastSplitHour: time.Now().Hour(),
	}

	logger.Init()
	return
}

func (f *FileLogger) Init(){
	// debug, trace, info 共用一个文件
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.file = file

	// warn, err, fatal 共用一个文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.warnFile = file

	go f.writeLog()
}

func (f *FileLogger) writeLog() {
	for logInfo := range f.LogInfoChan {
		if logInfo.IsWarn {
			if f.needSplitFile(f.warnFile) {
				f.splitFile(f.warnFile)
			}
			fmt.Fprintf(f.warnFile, logInfo.LogMsg)
		} else {
			if f.needSplitFile(f.file) {
				f.splitFile(f.file)
			}
			fmt.Fprintf(f.file, logInfo.LogMsg)
		}
	}
}

func (f *FileLogger) needSplitFile(file *os.File) (needSplit bool) {
	if f.logSplitType == LogSplitByHour {
		nowHour := time.Now().Hour()
		needSplit = f.lastSplitHour != nowHour
		if needSplit {
			f.lastSplitHour = nowHour
		}
	} else if f.logSplitType == LogSplitBySize {
		fileInfo, _ := file.Stat()
		needSplit = fileInfo.Size() >= f.logSplitSize
	}
	return
}

func (f *FileLogger) splitFile(file *os.File) {
	fileInfo, _ := file.Stat()
	fileName := fmt.Sprintf("%s/%s", f.logPath, fileInfo.Name())
	nowStr := time.Now().Format("2006_01_02_15:04:05")
	backupFileName := fmt.Sprintf("%s_%s",fileName, nowStr)

	file.Close()
	os.Rename(fileName, backupFileName)

	file, _ = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if strings.Contains(fileName, "wf") {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		f.level = LogLevelDebug
		return
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logInfo := GetlogInfo(LogLevelDebug, format, args...)
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logInfo := GetlogInfo(LogLevelTrace, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logInfo := GetlogInfo(LogLevelInfo, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logInfo := GetlogInfo(LogLevelWarn, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logInfo := GetlogInfo(LogLevelError, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logInfo := GetlogInfo(LogLevelFatal, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}