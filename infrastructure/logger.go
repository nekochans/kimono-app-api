package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOptions struct {
	LogLevel string
}

func newLogger(level zapcore.Level) (*zap.Logger, error) {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.TimeKey = "time"

	zapConfig := zap.Config{
		Level:            atomicLevel,
		Encoding:         "json",
		EncoderConfig:    cfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}

func NewLoggerFromOptions(o LoggerOptions) (*zap.Logger, error) {
	var levelEnabler zapcore.Level
	switch o.LogLevel {
	case "debug":
		levelEnabler = zap.DebugLevel
	case "warn":
		levelEnabler = zap.WarnLevel
	case "error":
		levelEnabler = zap.ErrorLevel
	default:
		levelEnabler = zap.InfoLevel
	}

	return newLogger(levelEnabler)
}
