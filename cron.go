package main

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"gin-blog/models"
)

func main() {
	log.Println("Starting...")

	//会根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	//会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息, 10秒后会往t1通道写内容 (当前时间)
	t1 := time.NewTimer(time.Second * 10)
	//阻塞 select 等待 channel
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
