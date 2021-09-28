package model

import "applyUpLoadFile/middleware/db"

type ApplyInfo struct {
	Id    int64  `json:"id"  xorm:"pk autoincr BIGINT(20)"`
	Name  string `json:"name"  xorm:"null default '' comment('申请人名称') VARCHAR(255)" `
	Phone string `json:"phone"  xorm:"null default '' comment('手机号') VARCHAR(255)" `
	Email string `json:"email"  xorm:"null default '' comment('邮箱') VARCHAR(255)" `
	CId   int64  `json:"c_id"  xorm:"INT(20)"`
}

func (u *ApplyInfo) TableName() string {
	return "apply_info"
}

func SelectApplyInfoById(id int64) (list []*ApplyInfo, err error) {
	db.WebDB.ShowSQL(true)
	err = db.WebDB.Where("id = ?", id).Find(&list)
	return
}

func SelectApplyInfo(ap *ApplyInfo) (apInfo ApplyInfo, err error) {
	db.WebDB.ShowSQL(true)
	_, err = db.WebDB.Where("name=? and phone=? and email=? and c_id=?",
		ap.Name,
		ap.Phone,
		ap.Email,
		ap.Id).
		Get(&apInfo)
	return
}
func ExistApplyInfo(ap *ApplyInfo) (exist bool, err error) {

	db.WebDB.ShowSQL(true)
	exist, err = db.WebDB.Where("name=? and phone=? and email=? and c_id=?",
		ap.Name,
		ap.Phone,
		ap.Email,
		ap.Id).
		Exist(ap)

	return
}

func InsertApplyInfo(applyInfo *ApplyInfo) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.Insert(applyInfo)
}

func DelApplyInfoById(applyInfo *ApplyInfo) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.ID(applyInfo.Id).Delete(applyInfo)
}

func UpdateApplyInfo(applyInfo *ApplyInfo) (int64, error) {
	db.WebDB.ShowSQL(true)
	return db.WebDB.ID(applyInfo.Id).Update(applyInfo)
}
