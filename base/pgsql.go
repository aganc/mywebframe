package base

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type PgSQLConf struct {
	Service         string        `yaml:"service"`
	DataBase        string        `yaml:"database"`
	Addr            string        `yaml:"addr"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Charset         string        `yaml:"charset"`
	MaxIdleConns    int           `yaml:"maxidleconns"`
	MaxOpenConns    int           `yaml:"maxopenconns"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
	ConnTimeOut     time.Duration `yaml:"connTimeOut"`
	WriteTimeOut    time.Duration `yaml:"writeTimeOut"`
	ReadTimeOut     time.Duration `yaml:"readTimeOut"`
}

func (conf *PgSQLConf) checkConf() {
	CommonSecretChange(prefix, *conf, conf)

	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 1000
	}
	if conf.ConnMaxLifeTime == 0 {
		conf.ConnMaxLifeTime = 3600 * time.Second
	}
	if conf.ConnTimeOut == 0 {
		conf.ConnTimeOut = 3 * time.Second
	}
	if conf.WriteTimeOut == 0 {
		conf.WriteTimeOut = 1 * time.Second
	}
	if conf.ReadTimeOut == 0 {
		conf.ReadTimeOut = 1 * time.Second
	}
}

func InitPgsqlClient(conf PgSQLConf) (client *gorm.DB, err error) {
	conf.checkConf()

	dsn := fmt.Sprintf("%s:%s@%s/%s?",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DataBase)

	if conf.Charset != "" {
		dsn = dsn + "&charset=" + conf.Charset
	}

	c := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		//Logger:                                   newLogger(&conf),
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	//dns1 := "postgres://vin:vin@localhost:20000/testdb"
	dsn = "postgres://" + dsn
	fmt.Println(dsn)
	client, err = gorm.Open(postgres.Open(dsn), c)
	if err != nil {
		return client, err
	}

	sqlDB, err := client.DB()
	if err != nil {
		return client, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	return client, nil
}
