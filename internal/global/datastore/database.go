package datastore

import (
	"fmt"
	"gin-temp/conf"
	cusLog "gin-temp/internal/global/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.Cfg.Db.User,
		conf.Cfg.Db.Password,
		conf.Cfg.Db.Host,
		conf.Cfg.Db.Port,
		conf.Cfg.Db.Table,
		conf.Cfg.Db.Charset,
		conf.Cfg.Db.ParseTime,
		conf.Cfg.Db.Loc,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dbLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Cfg.Db.Prefix,
			SingularTable: conf.Cfg.Db.SingularTable,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	cusLog.Logger.Info(fmt.Sprintf("数据库连接成功：%s:%d/%s", conf.Cfg.Db.User, conf.Cfg.Db.Port, conf.Cfg.Db.Table))
}
