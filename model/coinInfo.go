package model

import (
	"applyUpLoadFile/middleware/db"
	"time"
)

type CoinInfo struct {
	Id         int64     `json:"id"  xorm:"pk autoincr BIGINT(20)"`
	ChainId    int64     `json:"chain_id"  xorm:"BIGINT(20)"`
	Name       string    `json:"name"  xorm:"null default '' comment('币名称') VARCHAR(255)" `
	FullName   string    `json:"full_name" xorm:"null default '' comment('币全称名称') VARCHAR(255)" `
	CreateTime time.Time `json:"create_time"  xorm:" default '' comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time"  xorm:" default '' comment('创建时间') DATETIME"`
}

func (u *CoinInfo) TableName() string {
	return "coin_info"
}

func GetCoinById(c *CoinInfo, id int64) (*CoinInfo, error) {
	db.WebDB.ShowSQL(true)
	//_, err = db.WebDB.Get(ap)
	_, err := db.WebDB.Where("id=? ", id).Get(c)
	return c, err

}

func InsertCoinInfo(coinInfo *CoinInfo) (int64, error) {
	db.WebDB.ShowSQL(true)
	coinInfo.CreateTime = time.Now().UTC()
	return db.WebDB.Insert(coinInfo)
}

func FindCoinInfoAll() ([]CoinInfo, error) {
	cList := []CoinInfo{}
	db.WebDB.ShowSQL(true)
	err := db.WebDB.OrderBy("name").Find(&cList)
	if err != nil {
		return nil, err
	}
	return cList, nil
}
