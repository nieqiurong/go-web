package entity

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	Id          int64     `gorm:"primary_key;auto_increment"`
	CmbUid      string    `gorm:"type:varchar(50);index:cmb_uid;column:cmb_uid"`
	CmbEid      string    `gorm:"type:varchar(50);index:cmb_eid;column:cmb_eid"`
	WxId        string    `gorm:"type:varchar(50);index:wx_id;column:wx_id"`
	WxName      string    `gorm:"type:varchar(255);column:wx_name"`
	WxSex       string    `gorm:"type:varchar(255);column:wx_sex"`
	WxProvince  string    `gorm:"type:varchar(255);column:wx_province"`
	WxCity      string    `gorm:"type:varchar(255);column:wx_city"`
	LoginTime   time.Time `gorm:"type:datetime;column:login_time"`
	CreatedTime time.Time `gorm:"type:datetime;column:created_time"`
}

func (User) TableName() string {
	return "biz_user"
}

func SaveUser(CmbUid string, WxName string) (err error) {
	var user = User{}
	err = db.Where("cmb_uid = ?", CmbUid).Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user := User{CmbUid: CmbUid, WxName: WxName, LoginTime: time.Now(), CreatedTime: time.Now()}
			e := db.Create(&user).Error
			if e != nil {
				log.Println("save user fail ", e)
				return e
			}
			return nil
		}
		return err
	}
	user = User{Id: user.Id, LoginTime: time.Now()}
	err = db.Model(&user).Update(User{Id: user.Id, LoginTime: time.Now()}).Error
	if err != nil {
		log.Println("update user fail ", err)
		return err
	}
	return nil
}
