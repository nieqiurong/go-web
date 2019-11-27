package entity

import (
	"time"
)

type User struct {
	Id         int64     `xorm:"pk autoincr" json:"id"`
	Account    string    `xorm:"varchar(20) not null unique(un_index_account) account" json:"account"`
	PassWord   string    `xorm:"varchar(50) not null password" json:"pass_word"`
	Name       string    `xorm:"varchar(50) name" json:"name"`
	QQ         string    `xorm:"varchar(50) qq" json:"qq"`
	Email      string    `xorm:"varchar(50) email" json:"email"`
	LoginTime  time.Time `xorm:"datetime login_time default CURRENT_TIMESTAMP" json:"login_time"`
	CreateTime time.Time `xorm:"datetime create_time default CURRENT_TIMESTAMP" json:"create_time"`
}

func (User) TableName() string {
	return "t_user"
}

func FindUserByAccount(account string) (u *User, err error) {
	user := &User{}
	_, err = db.Where("account = ?", account).Get(user)
	return user, err
}
