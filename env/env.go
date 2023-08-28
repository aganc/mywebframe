package env

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const DefaultRootPath = "."

const EnvLogLevel = "LOG_LEVEL"

const (
	// 容器中的环境变量
	AXSClusterType = "AXS_CLUSTER_TYPE"
	DockAppName    = "APP_NAME"
	DockerRunEnv   = "RUN_ENV"
)

// RUN_ENV： (prod，tips，test)
const (
	RunEnvDEV = 0
	RunEnvQAT = 1
	RunEnvPRE = 2
	RunEnvPRD = 3
	RunEnvDOC = 4
)

var (
	LocalIP  string
	AppName  string
	RunMode  string
	ServerID string

	runEnv int

	rootPath        string
	dockerPlateForm bool

	logLevel string
)

func init() {
	LocalIP = "127.0.0.1"
	dockerPlateForm = false
	if r := os.Getenv(AXSClusterType); r != "" {
		dockerPlateForm = true
		// 容器里，appName在编排的时候决定
		if n := os.Getenv(DockAppName); n != "" {
			AppName = n
			println("docker env, APP_NAME=", n)
		} else {
			println("docker env, lack APP_NAME!!!")
		}
	}

	// 运行环境
	RunMode = gin.ReleaseMode
	r := os.Getenv(DockerRunEnv)
	switch r {
	case "doc":
		runEnv = RunEnvDOC
	case "prd":
		runEnv = RunEnvPRD
	case "pre":
		runEnv = RunEnvPRE
		RunMode = gin.DebugMode
	case "qat":
		runEnv = RunEnvQAT
		RunMode = gin.DebugMode
	case "dev":
		runEnv = RunEnvDEV
		RunMode = gin.DebugMode
	default:
		runEnv = RunEnvDEV
		RunMode = gin.DebugMode
	}

	gin.SetMode(RunMode)

	//initDBSecret()

	logLevel = os.Getenv(EnvLogLevel)
}

// 判断项目运行平台：容器 vs 开发环境
func IsDockerPlatform() bool {
	return dockerPlateForm
}

// 开发环境可手动指定SetAppName
func SetAppName(appName string) {
	if !dockerPlateForm {
		AppName = appName
	}
}

func GetAppName() string {
	return AppName
}

// SetRootPath 设置应用的根目录
func SetRootPath(r string) {
	if !dockerPlateForm {
		rootPath = r
	}
}

// RootPath 返回应用的根目录
func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	} else {
		return DefaultRootPath
	}
}

// GetConfDirPath 返回配置文件目录绝对地址
func GetConfDirPath() string {
	return filepath.Join(GetRootPath(), "conf")
}

// LogRootPath 返回log目录的绝对地址
func GetLogDirPath() string {
	return filepath.Join(GetRootPath(), "log")
}

func GetRunEnv() int {
	return runEnv
}

func SetServerID(serverID string) {
	ServerID = serverID
}

func GetServerID() string {
	return ServerID
}

func GetLogLevel() string {
	return logLevel
}
