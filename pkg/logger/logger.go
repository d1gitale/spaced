// Package logger defines logger using zap
package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func LoggerFromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey{}).(*zap.Logger); ok {
		return l
	}
	return zap.NewNop()
}

func NewLogger() *zap.Logger {
	l, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		log.Fatalf("failed to instantiate logger: %v", err)
	}
	return l
}
