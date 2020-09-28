package utils

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ugorm "ops.was.ink/opsmicro/user/basic/gorm"
	userconf "ops.was.ink/opsmicro/user/config"
	"time"
)

var (
	db *gorm.DB
	err error
)

func DBInit() error {
	// 1. 实例化 db 对象
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userconf.DBConf.Username,
		userconf.DBConf.Password,
		userconf.DBConf.Address,
		userconf.DBConf.DBName,
	)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return err
	}

	// 2. 开启调试模模式，可以打印出具体的 SQL 语句
	db = db.Debug()
	db.Use(&ugorm.GormTracePlugin{})

	// 3. 配置连接池
	sql_db, err := db.DB()
	if err != nil {
		return err
	}
	sql_db.SetMaxOpenConns(100)
	sql_db.SetMaxIdleConns(10)
	sql_db.SetConnMaxLifetime(time.Hour)
	// 4. 激活连接
	if err := sql_db.Ping(); err != nil {
		return err
	}

	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}