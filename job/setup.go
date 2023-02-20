package job

import (
	"github.com/robfig/cron/v3"
	"log"
)

func Setup() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("00 00 14 * * ?", func() { checkToken() })
	c.AddFunc("00 00 12 * * ? ", func() { GetBalance() })
	c.AddFunc("00 00 20 * * ? ", func() { GetBalance() })

	// 启动
	c.Start()
	log.Println("自动抢票中")

}
