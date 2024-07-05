package model

import (
	"fmt"
	"gin-rest-api/config"
	"gin-rest-api/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
	DeletedAt  int `json:"deleted_at"`
}

// Init 初始化数据库连接
func Init(cfg *config.App) error {
	v := cfg.Section("db")
	var (
		dsn         = v.Key("dsn").String()
		maxIdleConn = v.Key("max_idle_conn").MustInt(10)
		maxOpenConn = v.Key("max_open_conn").MustInt(30)
	)
	conn, err := openConn(dsn, maxIdleConn, maxOpenConn, cfg)
	if err != nil {
		return fmt.Errorf("open db conn failed, error: %s", err.Error())
	}
	db = conn
	return nil
}

func openConn(dsn string, idle, open int, cfg *config.App) (*gorm.DB, error) {
	var (
		openDB *gorm.DB
		err    error
	)
	if cfg.LogLevel == "debug" {
		newLogger := log.NewGormLogger(logger.Config{
			LogLevel: logger.Info,
		})
		openDB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: newLogger})
	} else {
		openDB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}
	db, err := openDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(open)
	return openDB, nil
}
