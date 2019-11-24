package entity

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-web/setting"
	"log"
	"time"
)

var db *xorm.Engine

func InitDb() {
	//username:password@protocol(address)/dbname?param=value
	dbConfig := setting.Application.Db
	user := dbConfig.UserName
	pwd := dbConfig.PassWord
	host := dbConfig.Host
	port := dbConfig.Port
	dbName := dbConfig.DbName
	var err error
	db, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, pwd, host, port, dbName))
	if err != nil {
		log.Fatal("connect database fail !", err)
	}
	duration, err := time.ParseDuration(dbConfig.MaxLifetime)
	if err != nil {
		log.Fatal("parseDuration fail !", err)
	}
	//db.ShowSQL(true)
	db.DB().SetConnMaxLifetime(duration)
	db.DB().SetMaxIdleConns(dbConfig.MaxIdle)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpen)
	//auto create table
	//_ = db.Sync2(&User{})
}
