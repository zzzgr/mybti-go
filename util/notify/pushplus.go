package notify

import (
	"log"
	"mybti-go/util/http"
)

func Pushplus(token string, title string, content string) {
	res, _ := http.NewRestyClient().R().
		SetFormData(map[string]string{
			"token":   token,
			"title":   title,
			"content": content,
		}).
		Execute("POST", "http://www.pushplus.plus/send")
	log.Println("pushplus通知结果: ", res)
}
