package db

import (
	conf "applyUpLoadFile/config"
	"applyUpLoadFile/middleware/log"
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var HooDB *MysqlDB

func init() {
	HooDB = InitFinanceDB(conf.Cfg.Databases["web"])

}

type MysqlDB struct {
	Name  string
	DB    *gorm.DB
	IDGen *snowflake.Node
}

// InitFinanceDB
// 连接数据库
func InitFinanceDB(cfg conf.Databases) *MysqlDB {

	dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.User, cfg.PassWord, cfg.Url, cfg.Name)
	db, err := InitDataBase(cfg.Type, dburl, cfg.Mode)
	if err != nil {
		panic(err.Error())
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
	return &MysqlDB{
		Name:  cfg.Name,
		DB:    db,
		IDGen: node,
	}
}

// InitDataBase
// 多数据库基础函数
func InitDataBase(dbType, dbUrl, dbMode string) (*gorm.DB, error) {
	if dbUrl == "" || dbType == "" {
		log.Infof("database's conf is null")
	}
	db, err := gorm.Open(dbType, dbUrl)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("gorm db is nil")
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(32)
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.LogMode(dbMode == "dev")
	return db, nil
}
