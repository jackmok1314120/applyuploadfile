package model

import (
	"applyUpLoadFile/middleware/db"
	"time"
)

type ApplyPending struct {
	Id              int64     `json:"id"  xorm:"pk autoincr BIGINT(20)"`
	Name            string    `json:"name"  xorm:"null default '' comment('申请人名称') VARCHAR(255)" `
	Phone           string    `json:"phone"  xorm:"null default '' comment('手机号') VARCHAR(255)" `
	Email           string    `json:"email"  xorm:"null default '' comment('邮箱') VARCHAR(255)" `
	CoinName        string    `json:"coin_name"  xorm:"null default '' comment('币种名') VARCHAR(255)" `
	Introduce       string    `json:"introduce"  xorm:"null default '' comment('介绍') VARCHAR(255)" `
	IdCardPicture   string    `json:"id_card_picture"  xorm:"null default '' comment('身份证复印件') TEXT" binding:"required" `
	BusinessPicture string    `json:"business_picture"  xorm:"null default '' comment('营业执照复印件') VARCHAR(255)" binding:"required"`
	Pass            int64     `json:"pass" xorm:" default 0 comment('营业执照复印件') INT(11)"`
	CreateTime      time.Time `json:"create_time"  xorm:" default '' comment('创建时间') DATETIME"`
	UpdateTime      time.Time `json:"update_time"  xorm:" default '' comment('创建时间') DATETIME"`
}

func (u *ApplyPending) TableName() string {
	return "apply_pending"
}

func ExistApply(ap *ApplyPending) (exist bool, err error) {

	db.WebDB.ShowSQL(true)
	//exist, err = db.WebDB.Exist(ap)
	exist, err = db.WebDB.Where("name=? and coin_name=? and phone=? and email=?",
		ap.Name,
		ap.CoinName,
		ap.Phone,
		ap.Email,
	).Exist(ap)
	return
}

func ExistApplyByPath(ap *ApplyPending) (exist bool, err error) {

	db.WebDB.ShowSQL(true)
	//exist, err = db.WebDB.Exist(ap)
	exist, err = db.WebDB.Where("id_card_picture=? and business_picture=?",
		ap.IdCardPicture,
		ap.BusinessPicture,
	).Exist(ap)
	return
}

func GetApply(ap *ApplyPending) (aps *ApplyPending, err error) {

	db.WebDB.ShowSQL(true)
	//_, err = db.WebDB.Get(ap)
	_, err = db.WebDB.Where("name=? and coin_name=? and phone=? and email=?",
		ap.Name,
		ap.CoinName,
		ap.Phone,
		ap.Email,
	).Get(ap)
	return ap, err
}

func InsertApplyPending(applyInfo *ApplyPending) (int64, error) {
	db.WebDB.ShowSQL(true)
	applyInfo.CreateTime = time.Now().UTC()
	applyInfo.Pass = 0
	return db.WebDB.Insert(applyInfo)
}

func UpdateApplyPending(applyInfo *ApplyPending) (int64, error) {

	db.WebDB.ShowSQL(true)
	applyInfo.UpdateTime = time.Now().UTC()
	return db.WebDB.ID(applyInfo.Id).Update(applyInfo)
}
