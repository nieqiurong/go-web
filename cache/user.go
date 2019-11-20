package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const cacheKey string = "user:cache"

type jsonTime time.Time

func (this jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type User struct {
	CmbUid      string   `json:"username"`
	WxName      string   `json:"nickname"`
	LoginTime   jsonTime `json:"login_time"`
	CreatedTime jsonTime `json:"created_time"`
}

func SaveUser(CmbUid string, WxName string) error {
	b, e := client.HExists(cacheKey, CmbUid).Result()
	if e != nil {
		log.Println("get user cache fail :", e)
		return e
	}
	if b {
		//预留更新
	} else {
		user := User{
			CmbUid:      CmbUid,
			WxName:      WxName,
			LoginTime:   jsonTime(time.Now()),
			CreatedTime: jsonTime(time.Now()),
		}
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("json.marshal failed, err:", err)
			return nil
		}
		client.HSet(cacheKey, CmbUid, data)
	}
	return nil
}
