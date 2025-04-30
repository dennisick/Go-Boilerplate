package ctx

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

type Key string

const RequestIDKey Key = "requestId"
const LoggerKey Key = "logger"
const RequestTimeKey Key = "requestTime"

func GetRequestID(ctx context.Context) string {
	id := ctx.Value(RequestIDKey).(string)
	return id
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

func GetLogger(ctx context.Context) zerolog.Logger {
	logger := ctx.Value(LoggerKey).(zerolog.Logger)
	return logger
}

func SetLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

func GetRequestTime(ctx context.Context) time.Time {
	return ctx.Value(RequestTimeKey).(time.Time)
}

func SetRequestTime(ctx context.Context, time time.Time) context.Context {
	return context.WithValue(ctx, RequestTimeKey, time)
}
