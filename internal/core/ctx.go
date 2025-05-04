package core

import (
	"context"

	"github.com/rs/zerolog"
)

type Key string

const LoggerKey Key = "logger"

func GetLogger(ctx context.Context) zerolog.Logger {
	logger := ctx.Value(LoggerKey).(zerolog.Logger)
	return logger
}

func SetLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}
