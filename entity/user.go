package entity

import (
	"time"
)

type User struct {
	Id         int64     `xorm:"pk autoincr"`
	Account    string    `xorm:"varchar(20) not null unique(un_index_account) account"`
	PassWord   string    `xorm:"varchar(50) not null password"`
	Name       string    `xorm:"varchar(50) name"`
	QQ         string    `xorm:"varchar(50) qq"`
	Email      string    `xorm:"varchar(50) email"`
	LoginTime  time.Time `xorm:"datetime login_time default CURRENT_TIMESTAMP"`
	CreateTime time.Time `xorm:"datetime create_time default CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "t_user"
}

func FindUserByAccount(account string) (u *User, err error) {
	user := &User{}
	_, err = db.Where("account = ?", account).Get(user)
	return user, err
}
