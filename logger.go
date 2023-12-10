package logger

import (
	"fmt"
	"runtime"
	"time"
)

type Logger struct {
	LogGroup string
}

func CreateLogger(logGroup string) *Logger {
	return &Logger{
		LogGroup: logGroup,
	}
}

func (logger *Logger) prepareMessage(message string) string {
	logGroup := logger.LogGroup
	if len(logGroup) <= 0 {
		logGroup = "default"
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("[%s, %s, %s:%d]\n%s", logGroup, now, truncateStart(filename, 20, "..."), line, message)
	}
	return fmt.Sprintf("[%s %s]\n%s", logGroup, now, message)
}

func (logger *Logger) Log(level LogLevel, message string, context ...interface{}) {
	Log(level, message, context...)
}

func (logger *Logger) Emergency(message string, context ...interface{}) {
	Emergency(logger.prepareMessage(message), context...)
}

func (logger *Logger) Alert(message string, context ...interface{}) {
	Alert(logger.prepareMessage(message), context...)
}

func (logger *Logger) Critical(message string, context ...interface{}) {
	Critical(logger.prepareMessage(message), context...)
}

func (logger *Logger) Error(message string, context ...interface{}) {
	Error(logger.prepareMessage(message), context...)
}

func (logger *Logger) Warning(message string, context ...interface{}) {
	Warning(logger.prepareMessage(message), context...)
}

func (logger *Logger) Notice(message string, context ...interface{}) {
	Notice(logger.prepareMessage(message), context...)
}

func (logger *Logger) Info(message string, context ...interface{}) {
	Info(logger.prepareMessage(message), context...)
}

func (logger *Logger) Debug(message string, context ...interface{}) {
	Debug(logger.prepareMessage(message), context...)
}
