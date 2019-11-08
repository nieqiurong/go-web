package schedule

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func InitSchedule() {
	c := cron.New()
	_ = c.AddFunc("@every 3s", func() {
		log.Println(time.Now().Format("2006/01/02 15:04:05"))
	})
	c.Start()
}
