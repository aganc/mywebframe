package helpers

import (
	"airport/base"
	"airport/conf"
	"github.com/smallnest/rpcx/log"
	"gorm.io/gorm"
)

var (
	MysqlClientNtExam *gorm.DB
)

func InitMysql() {
	var err error
	for name, dbConf := range conf.RConf.Mysql {
		switch name {
		case "base":
			MysqlClientNtExam, err = base.InitMysqlClient(dbConf)
		}
		if err != nil {
			//panic("mysql connect error: %v" + err.Error())
			log.Warnf("mysql connect error: %v" + err.Error())
		}
	}
}

func CloseMysql() {
	//_ = MysqlClientDemo.Close()
	//_ = MysqlClientTest.Close()
}
