package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web/setting"
	"log"
)

func InitDb() {
	//username:password@protocol(address)/dbname?param=value
	user := setting.Application.Db.UserName
	pwd := setting.Application.Db.PassWord
	host := setting.Application.Db.Host
	dbName := setting.Application.Db.DbName
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connect database fail !", err)
	}
	//auto create table
	db.AutoMigrate(&Student{})
}
