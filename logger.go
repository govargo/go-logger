package logger

import (
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error

	sn := os.Getenv("SERVICE_NAME")
	if sn == "" {
		config := NewProductionConfig()
		logger, err = config.Build(zap.AddCallerSkip(1))
	} else {
		logger, err = NewProductionWithCore(WrapCore(
			ReportAllErrors(true),
			ServiceName(sn),
		), zap.AddCallerSkip(1))
	}

	if err != nil {
		panic(err)
	}
}

// NewProductionWithCore is same as NewProduction but accepts a custom configured core
func NewProductionWithCore(core zap.Option, options ...zap.Option) (*zap.Logger, error) {
	options = append(options, core)

	return NewProductionConfig().Build(options...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func Sync() error {
	err := logger.Sync()
	if err != nil {
		return err
	}
	return nil
}
