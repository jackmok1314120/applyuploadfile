package model

import (
	"applyUpLoadFile/middleware/db"
	"time"
)

type ChainInfo struct {
	Id         int64     `json:"id"  xorm:"pk autoincr BIGINT(20)"`
	Name       string    `json:"name"  xorm:"null default '' comment('链名称') VARCHAR(255)" `
	CreateTime time.Time `json:"create_time"  xorm:" default '' comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time"  xorm:" default '' comment('创建时间') DATETIME"`
}

func (u *ChainInfo) TableName() string {
	return "chain_info"
}

func InsertChainInfo(chainInfo *ChainInfo) (int64, error) {
	db.WebDB.ShowSQL(true)
	chainInfo.CreateTime = time.Now().UTC()
	return db.WebDB.Insert(chainInfo)
}

func GetChainByName(name string) (*ChainInfo, error) {
	c := new(ChainInfo)
	db.WebDB.ShowSQL(true)
	//_, err = db.WebDB.Get(ap)
	_, err := db.WebDB.Where("name=? ", name).Get(c)
	return c, err

}
