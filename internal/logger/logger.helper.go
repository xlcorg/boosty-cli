package logger

import (
	"context"

	"go.uber.org/zap"
)

// context logger ----------------------------------------------------

type ctxKey struct {
}

var attachedLoggerKey = &ctxKey{}

// AttachToContext adds the current logger to the context with a variadic number of fields.
func AttachToContext(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, getLogger(ctx).With(fields...))
}

func Debug(ctx context.Context, msg string, kvs ...any) {
	getLogger(ctx).Debugw(msg, kvs...)
}

func ErrorKV(ctx context.Context, msg string, kvs ...any) {
	getLogger(ctx).Errorw(msg, kvs...)
}

func Error(ctx context.Context, err error, msg string, kvs ...any) {
	getLogger(ctx).Errorw(msg, append(kvs, "error", err)...)
}

func Info(ctx context.Context, msg string, kvs ...any) {
	getLogger(ctx).Infow(msg, kvs...)
}

func Warn(ctx context.Context, msg string, kvs ...any) {
	getLogger(ctx).Warnw(msg, kvs...)
}

func Fatal(ctx context.Context, err error, msg string) {
	getLogger(ctx).Fatalw(msg, "error", err)
}

func getLogger(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return globalLogger
}
