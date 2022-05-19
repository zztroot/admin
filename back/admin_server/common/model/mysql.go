package model

import (
	"context"
	"fmt"
	"log"

	"goods/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbMap = map[string]*gorm.DB{}

type MysqlConn struct {
	db *gorm.DB
}

func (m *MysqlConn) GetMysqlDB(ctx context.Context, name ...string) *gorm.DB {
	if len(name) > 0 && name[0] != "default" {
		m.db = dbMap[name[0]]
	}
	if m.db == nil {
		key := "default"
		if len(name) > 0 {
			key = name[0]
		}
		m.db = dbMap[key]
	}
	return m.db
}

func InitMysql() {
	// 获取配置文件
	for name, c := range config.GlobalConfig.AppConfig.Database {
		conn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, c.User, c.Password, c.Host, c.Port, c.Database)
		// 连接数据库
		db, err := gorm.Open(c.Dialect, conn)
		if err != nil {
			log.Fatalln("connect error of database mysql:", err)
		}
		db.Set("", "ENGINE=InnoDB").AutoMigrate(GetObjects()...)
		db.DB().SetMaxIdleConns(c.MaxIdleConnNum)
		db.DB().SetMaxOpenConns(c.MaxOpenConnNum)
		dbMap[name] = db
	}
}
