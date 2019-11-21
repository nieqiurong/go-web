package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id         int64     `gorm:"primary_key;auto_increment"`
	Account    string    `gorm:"type:varchar(20);not null;unique_index:un_index_account;column:account"`
	PassWord   string    `gorm:"type:varchar(50);column:password"`
	Name       string    `gorm:"type:varchar(50);column:name"`
	QQ         string    `gorm:"type:varchar(50);column:qq"`
	Email      string    `gorm:"type:varchar(50);column:email"`
	LoginTime  time.Time `gorm:"type:datetime;column:login_time"`
	CreateTime time.Time `gorm:"type:datetime;column:create_time"`
}

func (User) TableName() string {
	return "t_user"
}

func FindUserByAccount(account string) (u *User, err error) {
	user := &User{}
	err = db.Where("account = ?", account).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, nil
}
