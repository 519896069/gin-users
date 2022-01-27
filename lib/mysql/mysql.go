package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	appConfig "user/config"
)

var Mysql *MysqlConnect

type MysqlConnect struct {
	Db *gorm.DB
}

func init() {
	database := appConfig.CONFIG.Database
	//加载mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username, database.Password, database.Host, database.Port, database.Database,
	)
	db, mysqlConnectErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if mysqlConnectErr != nil {
		panic(mysqlConnectErr)
	}
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//db.AutoMigrate(&models.User{}, &models.Token{})
	Mysql = &MysqlConnect{
		db,
	}
}
