package zlog

import (
	"airport/env"
	"github.com/smallnest/rpcx/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 对用户暴露的log配置
type LograteConf struct {
	MaxAge     int  `yaml:"max_age"`
	MaxBackups uint `yaml:"max_backups"`
	MaxSize    uint `yaml:"max_size"`
}
type LogConfig struct {
	Level       string      `yaml:"level"`
	Stdout      bool        `yaml:"stdout"`
	Lograte     bool        `yaml:"lograte"`
	LograteConf LograteConf `yaml:"lograte_conf"`
}

type loggerConfig struct {
	ZapLevel zapcore.Level

	// 以下变量仅对开发环境生效
	Stdout   bool
	Log2File bool
	Path     string

	// 日志切割配置
	Lograte    bool
	MaxAge     int
	MaxBackups uint
	MaxSize    uint
}

// 全局配置 仅限Init函数进行变更
var logConfig = loggerConfig{
	ZapLevel: zapcore.InfoLevel,

	Stdout:   false,
	Log2File: true,
	Path:     "./log",

	Lograte:    false,
	MaxAge:     0,
	MaxBackups: 0,
}

func InitLog(conf LogConfig) *zap.SugaredLogger {
	if err := RegisterAXSJSONEncoder(); err != nil {
		panic(err)
	}

	if level := env.GetLogLevel(); level != "" {
		conf.Level = level
	}

	logConfig.ZapLevel = getLogLevel(conf.Level)
	if env.IsDockerPlatform() {
		// 容器环境
		logConfig.Log2File = false
		logConfig.Stdout = true
	} else {
		// 开发环境下默认输出到文件，支持自定义是否输出到终端
		logConfig.Log2File = true
		logConfig.Stdout = conf.Stdout
		logConfig.Path = env.GetLogDirPath()
	}

	if conf.Lograte {
		logConfig.Lograte = conf.Lograte
		if conf.LograteConf.MaxAge > 0 && conf.LograteConf.MaxBackups > 0 {
			panic("max_age and max_backups cannot be both set")
		}
		if conf.LograteConf.MaxAge > 0 && conf.LograteConf.MaxSize > 0 {
			panic("max_age and max_size cannot be both set")
		}
		if conf.LograteConf.MaxAge > 0 {
			logConfig.MaxAge = conf.LograteConf.MaxAge
		}
		if conf.LograteConf.MaxBackups >= 0 {
			logConfig.MaxBackups = conf.LograteConf.MaxBackups
		}
		if conf.LograteConf.MaxSize >= 0 {
			logConfig.MaxSize = conf.LograteConf.MaxSize
		}
	}

	SugaredLogger = GetLogger()
	log.SetLogger(&rpcxLogger{})
	return SugaredLogger
}
