package helpers

import (
	"airport/base"
	"airport/conf"
	"github.com/smallnest/rpcx/log"
	"gorm.io/gorm"
)

var (
	PgsqlClientNtExam *gorm.DB
)

func InitPgsql() {
	var err error
	for name, dbConf := range conf.RConf.PgSQL {
		switch name {
		case "base":
			PgsqlClientNtExam, err = base.InitPgsqlClient(dbConf)
		}
		if err != nil {
			log.Warnf("pgsql connect error: %v" + err.Error())
		}
	}
}
