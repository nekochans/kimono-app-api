package infrastructure

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestCreateNewLoggerLogLevelIsDebug(t *testing.T) {
	o := &LoggerOptions{LogLevel: "debug"}
	l, err := NewLoggerFromOptions(*o)
	if err != nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsDebug Fatal", err)
	}

	ex := "TestCreateNewLoggerLogLevelIsDebug Message🐱"

	c := l.Check(zapcore.DebugLevel, ex)
	if c == nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsDebug Fatal", err)
	}

	// LogLevel: debug なのでdebugログが出力される事が期待される
	if c.Message != ex {
		t.Error("\nActually: ", c.Message, "\nExpected: ", ex)
	}
}

func TestCreateNewLoggerLogLevelIsWarn(t *testing.T) {
	o := &LoggerOptions{LogLevel: "warn"}
	l, err := NewLoggerFromOptions(*o)
	if err != nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsWarn Fatal", err)
	}

	// LogLevel: warn なのでdebugログは出力されない事が期待される
	c := l.Check(zapcore.DebugLevel, "TestCreateNewLoggerLogLevelIsWarn Message")
	if c != nil {
		t.Error("\nActually: ", c, "\nExpected: ", nil)
	}
}

func TestCreateNewLoggerLogLevelIsError(t *testing.T) {
	o := &LoggerOptions{LogLevel: "error"}
	l, err := NewLoggerFromOptions(*o)
	if err != nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsError Fatal", err)
	}

	// LogLevel: error なのでwarnログは出力されない事が期待される
	c := l.Check(zapcore.WarnLevel, "TestCreateNewLoggerLogLevelIsError Message")
	if c != nil {
		t.Error("\nActually: ", c, "\nExpected: ", nil)
	}
}

func TestCreateNewLoggerLogLevelIsInfo(t *testing.T) {
	o := &LoggerOptions{LogLevel: "info"}
	l, err := NewLoggerFromOptions(*o)
	if err != nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsDebug Fatal", err)
	}

	ex := "TestCreateNewLoggerLogLevelIsInfo Message🐱"

	c := l.Check(zapcore.InfoLevel, ex)
	if c == nil {
		t.Fatal("TestCreateNewLoggerLogLevelIsInfo Fatal", err)
	}

	// LogLevel: info なのでinfoログが出力される事が期待される
	if c.Message != ex {
		t.Error("\nActually: ", c.Message, "\nExpected: ", ex)
	}
}
