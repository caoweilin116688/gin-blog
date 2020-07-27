package main

import (
	"fmt"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"syscall"

	"gin-blog/models"
	"gin-blog/pkg/logging"

	"github.com/fvbock/endless"
)

func main() {
	setting.Setup() //加载app.ini 配置
	models.Setup()  //初始化DB
	logging.Setup() //初始化日志
	gredis.Setup()  //初始化redis

	//优雅重启
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	//非优雅重启
	/*router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()*/
}
