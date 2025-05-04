package api

import (
	"context"
	"time"
)

type Key string

const RequestIDKey Key = "requestId"
const RequestTimeKey Key = "requestTime"

func GetRequestID(ctx context.Context) string {
	id := ctx.Value(RequestIDKey).(string)
	return id
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

func GetRequestTime(ctx context.Context) time.Time {
	return ctx.Value(RequestTimeKey).(time.Time)
}

func SetRequestTime(ctx context.Context, time time.Time) context.Context {
	return context.WithValue(ctx, RequestTimeKey, time)
}
