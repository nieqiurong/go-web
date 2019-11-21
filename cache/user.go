package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	cacheKey   string = "user:cache"
	dateLayOut        = "2006-01-02 15:04:05"
)

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(dateLayOut))
	return []byte(stamp), nil
}

func (t *jsonTime) UnmarshalJSON(value []byte) error {
	dateStr := string(value)
	if dateStr != "" && dateStr != "null" {
		d, err := time.Parse(`"`+dateLayOut+`"`, dateStr)
		if err != nil {
			log.Println("json.unmarshalJSON failed , err :", err)
			return err
		}
		*t = jsonTime(d)
		return nil
	}
	return nil
}

type User struct {
	CmbUid      string   `json:"username"`
	WxName      string   `json:"nickname"`
	LoginTime   jsonTime `json:"login_time"`
	CreatedTime jsonTime `json:"created_time"`
}

func SaveUser(CmbUid string, WxName string) error {
	// 这个方法只是用来模拟存数据库.
	b, err := client.HExists(cacheKey, CmbUid).Result()
	var user User
	if err != nil {
		log.Println("client.hexists failed , err :", err)
		return err
	}
	if b {
		v, err := client.HGet(cacheKey, CmbUid).Result()
		if err != nil {
			log.Println("client.hget failed , err :", err)
			return err
		}
		err = json.Unmarshal([]byte(v), &user)
		if err != nil {
			fmt.Println("json.unmarshal failed, err:", err)
			return err
		}
	} else {
		user = User{
			CmbUid:      CmbUid,
			WxName:      WxName,
			LoginTime:   jsonTime(time.Now()),
			CreatedTime: jsonTime(time.Now()),
		}
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return err
	}
	client.HSet(cacheKey, CmbUid, data)
	return nil
}
