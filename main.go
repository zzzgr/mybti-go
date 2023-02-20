package main

import (
	"fmt"
	"log"
	"mybti-go/job"
	"mybti-go/setting"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("配置文件加载失败, 请在文件同级目录添加配置文件, err: %v\n", err)
		return
	}

	fmt.Println(`
                 _     _   _                   
 _ __ ___  _   _| |__ | |_(_)       __ _  ___  
| '_ ` + "`" + ` _ \| | | | '_ \| __| |_____ / _` + "`" + ` |/ _ \ 
| | | | | | |_| | |_) | |_| |_____| (_| | (_) |
|_| |_| |_|\__, |_.__/ \__|_|      \__, |\___/`)

	// 设置定时任务
	go func() {
		job.Setup()
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置5秒超时
	quit := make(chan os.Signal, 1)                      // 创建一个接受信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	log.Println("自动抢票已停止")
}
