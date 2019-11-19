package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web/setting"
	"log"
	"time"
)

var db *gorm.DB

func InitDb() {
	//username:password@protocol(address)/dbname?param=value
	dbConfig := setting.Application.Db
	user := dbConfig.UserName
	pwd := dbConfig.PassWord
	host := dbConfig.Host
	port := dbConfig.Port
	dbName := dbConfig.DbName
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, pwd, host, port, dbName))
	if err != nil {
		log.Fatal("connect database fail !", err)
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	//auto create table
	db.AutoMigrate(&Student{})
}
