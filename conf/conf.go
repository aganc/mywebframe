package conf

import (
	"airport/base"
	"airport/zlog"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"path/filepath"
	"time"
)

// 配置文件对应的全局变量
var (
	BasicConf TBasic
	Api       TApi
	RConf     ResourceConf
)

type ServerConfig struct {
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readtimeout"`
	WriteTimeout time.Duration `yaml:"writetimeout"`
}

// TBasic 基础配置,对应config.yaml
type TBasic struct {
	Log  zlog.LogConfig
	HTTP ServerConfig
	// ....业务可扩展其他简单的配置
}

// TApi 对应api.yaml
type TApi struct {
}

// ResourceConf 资源配置

// ResourceConf 对应resource.yaml
type ResourceConf struct {
	DBPrefix string `yaml:"db_prefix"`
	Mysql    map[string]base.MysqlConf
	PgSQL    map[string]base.PgSQLConf
	Redis    map[string]base.RedisConf
	//ClickHouse map[string]base.ClickHouseConf
	//HBase      map[string]hbase.HBaseConf
	//Elastic    map[string]base.ElasticClientConfig
	//Rmq        map[string]rmq.ClientConfig
	//KafkaPub   map[string]base.KafkaProducerConfig
	//KafkaSub   map[string]command.KafkaConsumeConfig
	//AOS        map[string]aos.Client
}

func InitConf() {
	// 加载通用基础配置（必须）
	LoadConf("config.yaml", SubConfMount, &BasicConf)

	// 加载资源类配置（optional）
	LoadConf("resource.yaml", SubConfMount, &RConf)

	// 加载api调用相关配置（optional）
	//env.LoadConf("api.yaml", env.SubConfMount, &Api)

	// 加载业务类(需要通过配置中心可修改的业务类配置)配置 （optional）
	// ... 加载更多配置
	//LoadConf("custom.yaml", SubConfMount, &Custom)
}

const (
	SubConfDefault = ""
	SubConfMount   = "mount"
	SubConfApp     = "app"
)

func LoadConf(filename, subConf string, s interface{}) {
	var path string
	path = filepath.Join(GetConfDirPath(), subConf, filename)

	if yamlFile, err := ioutil.ReadFile(path); err != nil {
		panic(filename + " get error: %v " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: %v" + err.Error())
	}
}

// GetRootPath 返回应用的根目录
func GetRootPath() string {
	return SubConfDefault
}

// GetConfDirPath 返回配置文件目录绝对地址
func GetConfDirPath() string {
	return filepath.Join(GetRootPath(), "conf")
}
