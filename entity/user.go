package entity

import (
	"log"
	"time"
)

type User struct {
	Id          int64     `xorm:"pk autoincr"`
	CmbUid      string    `xorm:"varchar(50) index(index_cmb_uid) cmb_uid"`
	CmbEid      string    `xorm:"varchar(50) index(index_cmb_eid) cmb_eid"`
	WxId        string    `xorm:"varchar(50) index(index_wx_id) wx_id"`
	WxName      string    `xorm:"varchar(20) wx_name"`
	WxSex       string    `xorm:"varchar(10) wx_sex"`
	WxProvince  string    `xorm:"varchar(10) wx_province"`
	WxCity      string    `xorm:"varchar(10) wx_city"`
	LoginTime   time.Time `xorm:"datetime login_time"`
	CreatedTime time.Time `xorm:"datetime created_time"`
}

func (User) TableName() string {
	return "biz_user"
}

func SaveUser(CmbUid string, WxName string) (err error) {
	var user = User{}
	session := db.NewSession()
	defer session.Close()
	//b, err := session.Where("cmb_uid = ?", CmbUid).Get(&user)
	bytes, err := session.Query("select * from biz_user where cmb_uid = ?", CmbUid)
	if err != nil {
		log.Println("select user fail ", err)
		return err
	}
	if bytes != nil {
		// 更新的不用管,不参与对比
		_, err = session.Update(User{Id: user.Id, LoginTime: time.Now()})
		if err != nil {
			log.Println("update user fail ", err)
			return err
		}
		return nil
	} else {
		//user := User{CmbUid: CmbUid, WxName: WxName, LoginTime: time.Now(), CreatedTime: time.Now()}
		//_, e := session.Insert(&user)
		_, e := session.Exec("insert into biz_user(cmb_uid,wx_name,login_time,created_time) VALUE(?,?,now(),now())", CmbUid, WxName)
		if e != nil {
			log.Println("save user fail ", e)
			return e
		}
		return nil
	}
}
