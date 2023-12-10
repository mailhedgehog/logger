package logger

import (
	"bytes"
	"fmt"
	"github.com/mailhedgehog/gounit"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

var capturedOutput string

func TestDebug(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(DebugColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestDebugViaLogGeneralHandler(t *testing.T) {
	MinLogLevel = DEBUG

	LogHandler = func(message string, context ...interface{}) {
		fmt.Println("MSG: " + message)
	}

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("MSG: Foo\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)

	LogHandler = nil
	LogDebugHandler = nil
}

func TestDebugViaLogDebugHandler(t *testing.T) {
	MinLogLevel = DEBUG

	LogDebugHandler = func(message string, context ...interface{}) {
		fmt.Println("Custom: " + message)
	}

	LogHandler = func(message string, context ...interface{}) {
		fmt.Println("MSG: " + message)
	}

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("Custom: Foo\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Debug("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)

	LogHandler = nil
	LogDebugHandler = nil
}

func TestInfo(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Info("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(InfoColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Info("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestNotice(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Notice("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(NoticeColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Notice("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestWarning(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Warning("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(WarningColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Warning("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestError(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Error("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(ErrorColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Error("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestCritical(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Critical("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(CriticalColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Critical("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestAlert(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Alert("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(AlertColor("Foo")+"\n", capturedOutput)

	MinLogLevel = EMERGENCY

	capturedOutput = captureOutput(func() {
		Alert("Foo")
	})
	(*gounit.T)(t).AssertEqualsString("", capturedOutput)
}

func TestEmergency(t *testing.T) {
	MinLogLevel = DEBUG

	capturedOutput = captureOutput(func() {
		Emergency("Foo")
	})
	(*gounit.T)(t).AssertEqualsString(EmergencyColor("Foo")+"\n", capturedOutput)
}
