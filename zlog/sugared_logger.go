package zlog

import (
	"context"

	"airport/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetLogger 获得一个新的logger 会把日志打印到 name.log 中，不建议业务使用
// deprecated
func GetLogger() (s *zap.SugaredLogger) {
	if SugaredLogger == nil {
		SugaredLogger = newLogger().WithOptions(zap.AddCallerSkip(1)).Sugar()
	}
	return SugaredLogger
}

// 通用字段封装
func sugaredLogger(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return SugaredLogger
	}

	return SugaredLogger.With(
		//zap.String("logId", GetLogID(ctx)),
		//zap.String("requestId", GetRequestID(ctx)),
		zap.String("module", env.GetAppName()),
		zap.String("localIp", env.LocalIP),
		//zap.String("uri", GetKeyURI(ctx)),
	)
}

// 提供给业务使用的server log 日志打印方法
func Debug(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.DebugLevel) {
		return
	}

	sugaredLogger(ctx).Debug(args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.DebugLevel) {
		return
	}
	sugaredLogger(ctx).Debugf(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.InfoLevel) {
		return
	}
	sugaredLogger(ctx).Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.InfoLevel) {
		return
	}
	sugaredLogger(ctx).Infof(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.WarnLevel) {
		return
	}
	sugaredLogger(ctx).Warn(args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.WarnLevel) {
		return
	}
	sugaredLogger(ctx).Warnf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.ErrorLevel) {
		return
	}
	sugaredLogger(ctx).Error(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.ErrorLevel) {
		return
	}
	sugaredLogger(ctx).Errorf(format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.PanicLevel) {
		return
	}
	sugaredLogger(ctx).Panic(args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.PanicLevel) {
		return
	}
	sugaredLogger(ctx).Panicf(format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	if NoLog(ctx, zapcore.FatalLevel) {
		return
	}
	sugaredLogger(ctx).Fatal(args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	if NoLog(ctx, zapcore.FatalLevel) {
		return
	}
	sugaredLogger(ctx).Fatalf(format, args...)
}
