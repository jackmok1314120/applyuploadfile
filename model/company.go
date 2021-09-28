package model

import "applyUpLoadFile/middleware/db"

type Company struct {
	Id              int64  `json:"id"  xorm:"pk autoincr INT(20)"`
	CoinName        string `json:"coin_name"  xorm:"null default '' comment('币种名') VARCHAR(255)" `
	Introduce       string `json:"introduce"  xorm:"null default '' comment('介绍') VARCHAR(255)" `
	IdCardPicture   string `json:"id_card_picture"  xorm:"null default '' comment('身份证复印件') TEXT" binding:"required" `
	BusinessPicture string `json:"business_picture"  xorm:"null default '' comment('营业执照复印件') VARCHAR(255)" binding:"required" `
}

func (u *Company) TableName() string {
	return "company"
}

func SelectCompanyById(id int64) (list []*Company, err error) {
	db.WebDB.ShowSQL(true)
	err = db.WebDB.Where("id = ?", id).Find(&list)
	return
}

func ExistCompany(ap *Company) (exist bool, err error) {

	db.WebDB.ShowSQL(true)
	exist, err = db.WebDB.Where("apply_id=?", ap.Id).Exist(ap)

	return
}

func InsertCompany(companyInfo *Company) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.Insert(companyInfo)
}

func DelCompanyById(companyInfo *Company) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.ID(companyInfo.Id).Delete(companyInfo)
}

func UpdateCompany(companyInfo *Company) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.ID(companyInfo.Id).Update(companyInfo)
}
