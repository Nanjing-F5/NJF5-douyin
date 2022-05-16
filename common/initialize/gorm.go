package initialize

import (
	"douyin/global"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Mysql配置MySQL数据库
func Mysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}

	// 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间
	sqlDb.SetConnMaxLifetime(time.Hour)

	global.Db = db
}
