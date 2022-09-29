package log

import (
	"go.uber.org/zap"
)

type Logger interface {
	Error(msg string)
}

var ErrorLogger Logger

func SetLogger(logger Logger) {
	ErrorLogger = logger
}

type zapLogger struct {
	logger *zap.Logger
}

func (zl *zapLogger) Error(msg string) {
	zl.logger.Error(msg)
}

func newZapLogger() *zapLogger {
	z, _ := zap.NewDevelopment()
	return &zapLogger{logger: z}
}

func init() {
	SetLogger(newZapLogger())
}
