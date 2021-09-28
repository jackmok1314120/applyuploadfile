package db

import (
	conf "applyUpLoadFile/config"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var WebDB *xorm.Engine

func init() {
	WebDB = InitConnDB2(conf.Cfg.Databases["web"])
}

// InitConnDB2
// 连接数据库
func InitConnDB2(cfg conf.Databases) *xorm.Engine {
	dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.User, cfg.PassWord, cfg.Url, cfg.Name)
	conn, e := initDBConn(cfg.Type, dburl, cfg.Mode)
	if e != nil {
		panic(e.Error())
	}
	return conn
}

// initDBConn
// 初始化数据库连接
func initDBConn(dbType, dbUrl, dbMode string) (coin *xorm.Engine, err error) {
	if dbUrl == "" || dbType == "" {
		return nil, errors.New("empty databases config")
	}
	conn, err := xorm.NewEngine(dbType, dbUrl)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(6)
	conn.SetConnMaxLifetime(60 * time.Second)
	//conn.ShowSQL(true)
	//conn.ShowExecTime(true)
	return conn, nil
}
