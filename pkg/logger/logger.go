// Package logger defines logger using zap
package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type loggerKey struct{}

type Logger struct {
	l *zap.Logger
}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func LoggerFromCtx(ctx context.Context) *Logger {
	if l, ok := ctx.Value(loggerKey{}).(*zap.Logger); ok {
		return &Logger{l: l}
	}
	return &Logger{l: zap.NewNop()}
}

func NewLogger() *Logger {
	l, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		log.Fatalf("failed to instantiate logger: %v", err)
	}
	return &Logger{l: l}
}

func (l *Logger) Fatal(msg string, err error) {
	l.l.Fatal(msg, zap.Error(err))
}

func (l *Logger) Error(msg string, err error) {
	l.l.Error(msg, zap.Error(err))
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}
