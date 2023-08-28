package zlog

import (
	"airport/env"
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetZapLogger() (l *zap.Logger) {
	if ZapLogger == nil {
		ZapLogger = newLogger().WithOptions(zap.AddCallerSkip(1))
	}
	return ZapLogger
}

func zapLogger(ctx context.Context) *zap.Logger {
	m := GetZapLogger()
	//m = m.WithOptions(zap.AddCallerSkip(1))
	if ctx == nil {
		return m
	}
	return m.With(
		zap.String("logId", GetLogID(ctx)),
		zap.String("requestId", GetRequestID(ctx)),
		zap.String("module", env.GetAppName()),
		zap.String("localIp", env.LocalIP),
		zap.String("uri", GetKeyURI(ctx)),
	)
}

func DebugLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.DebugLevel) {
		return
	}
	zapLogger(ctx).Debug(msg, fields...)
}
func InfoLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.InfoLevel) {
		return
	}
	zapLogger(ctx).Info(msg, fields...)
}

func WarnLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.WarnLevel) {
		return
	}
	zapLogger(ctx).Warn(msg, fields...)
}

func ErrorLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.ErrorLevel) {
		return
	}
	zapLogger(ctx).Error(msg, fields...)
}

func PanicLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.PanicLevel) {
		return
	}
	zapLogger(ctx).Panic(msg, fields...)
}

func FatalLogger(ctx context.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx, zapcore.FatalLevel) {
		return
	}
	zapLogger(ctx).Fatal(msg, fields...)
}
