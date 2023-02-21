package notify

import (
	"fmt"
	"log"
	"mybti-go/util/http"
	"time"
)

func Pushplus(token string, title string, content string) {
	content = fmt.Sprintf("%s\n\n 当前时间: %s", content, time.Now().Format("2006-01-02 15:04:05"))
	res, _ := http.NewRestyClient().R().
		SetFormData(map[string]string{
			"token":   token,
			"title":   title,
			"content": content,
		}).
		Execute("POST", "http://www.pushplus.plus/send")
	log.Println("pushplus通知结果: ", res)
}
