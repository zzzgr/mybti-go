package job

import (
	"encoding/base64"
	"fmt"
	"log"
	"mybti-go/setting"
	"mybti-go/util/notify"
	"strconv"
	"strings"
	"time"
)

func checkToken() {
	for _, u := range setting.GlobalConfig.User {
		// decode
		bytes, err := base64.StdEncoding.DecodeString(u.AccessToken)
		if err != nil {
			log.Println("")
			notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken,
				"token有效期通知",
				fmt.Sprintf("%s token格式配置有误\n", u.Name),
			)
			return
		}
		// 获取token中的时间戳
		str := string(bytes)
		ts, _ := strconv.Atoi(strings.Split(str, ",")[1])
		ttl := time.Unix(int64(ts/1000), 0).Sub(time.Now())
		hour := ttl.Hours()

		// 通知
		var msg string
		if hour > 24 {
			msg = fmt.Sprintf("%s 的token还有%d天失效", u.Name, (int)(hour/24))
		} else {
			msg = fmt.Sprintf("%s 的token还有%f小时就要到期了，请赶快更新", u.Name, hour)
		}
		notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken,
			"token有效期通知", msg)
	}
}
