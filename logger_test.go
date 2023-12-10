package logger

import (
	"fmt"
	"github.com/mailhedgehog/gounit"
	"testing"
	"time"
)

func TestLoggerDebug(t *testing.T) {
	logManager := CreateLogger("foo")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(DebugColor(fmt.Sprintf("[foo, %s, ...ogger/logger_test.go:17]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerInfo(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Info("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(InfoColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:36]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Info("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerNotice(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Notice("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(NoticeColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:55]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Notice("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerWarning(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Warning("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(WarningColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:74]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Warning("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerError(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Error("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(ErrorColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:93]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Error("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerCritical(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Critical("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(CriticalColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:112]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Critical("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerAlert(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Alert("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(AlertColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:131]\nFoo", now))+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		logManager.Alert("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestLoggerEmergency(t *testing.T) {
	logManager := CreateLogger("bar")
	now := time.Now().Format("2006-01-02 15:04:05")

	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		logManager.Emergency("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(EmergencyColor(fmt.Sprintf("[bar, %s, ...ogger/logger_test.go:150]\nFoo", now))+"\n", capturedOutput)
}
