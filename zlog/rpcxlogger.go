package zlog

type rpcxLogger struct {
}

func (l *rpcxLogger) Debug(v ...interface{}) {
	sugaredLogger(nil).Debug(v)
}

func (l *rpcxLogger) Debugf(format string, v ...interface{}) {
	sugaredLogger(nil).Debugf(format, v)
}

func (l *rpcxLogger) Info(v ...interface{}) {
	sugaredLogger(nil).Info(v)
}

func (l *rpcxLogger) Infof(format string, v ...interface{}) {
	sugaredLogger(nil).Infof(format, v)
}

func (l *rpcxLogger) Warn(v ...interface{}) {
	sugaredLogger(nil).Warn(v)
}

func (l *rpcxLogger) Warnf(format string, v ...interface{}) {
	sugaredLogger(nil).Warnf(format, v)
}

func (l *rpcxLogger) Error(v ...interface{}) {
	sugaredLogger(nil).Error(v)
}

func (l *rpcxLogger) Errorf(format string, v ...interface{}) {
	sugaredLogger(nil).Errorf(format, v)
}

func (l *rpcxLogger) Fatal(v ...interface{}) {
	sugaredLogger(nil).Fatal(v)
}

func (l *rpcxLogger) Fatalf(format string, v ...interface{}) {
	sugaredLogger(nil).Fatalf(format, v)
}

func (l *rpcxLogger) Panic(v ...interface{}) {
	sugaredLogger(nil).Panic(v)
}

func (l *rpcxLogger) Panicf(format string, v ...interface{}) {
	sugaredLogger(nil).Panicf(format, v)
}
