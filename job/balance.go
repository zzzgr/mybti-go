package job

import (
	"fmt"
	"log"
	"mybti-go/job/dto"
	"mybti-go/setting"
	http "mybti-go/util/http"
	"mybti-go/util/notify"
	"strconv"
	"time"
)

func GetBalance() {
	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowStr := tomorrow.Format("20060102")
	res, _ := http.NewRestyClient().R().
		Execute("GET", "https://tool.bitefu.net/jiari/?d="+tomorrowStr)
	isHoliday, _ := strconv.Atoi(string(res.Body()))

	if isHoliday == 0 {
		// 抢票
		log.Printf("%s, 抢票\n", tomorrowStr)
		doGetBalance(&tomorrow)
	} else {
		// 不抢票
		log.Printf("%s, 不抢票\n", tomorrowStr)
		notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken,
			"放假通知",
			fmt.Sprintf("明天(%s)放假一天", tomorrowStr),
		)
	}
}

func doGetBalance(date *time.Time) {
	tomorrowStr := date.Format("20060102")
	for _, u := range setting.GlobalConfig.User {
		var title, msg string
		balance, count, err := singleGetBalance(&u, date)
		if err != nil {
			title = "抢票异常通知"
			msg = fmt.Sprintf("%s抢的%s得票失败\n 错误信息: %s \n", u.Name, tomorrowStr, err.Error())
			notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken, title, msg)
			continue
		}

		if balance.Balance > 0 {
			// 成功
			msg = fmt.Sprintf("%s抢票成功 \n 线路: %s-%s \n 进站时间: %s \n 进站口: %s \n 抢票次数: %d",
				u.Name, u.Line, u.Station, u.Time, balance.StationEntrance, count)
			title = fmt.Sprintf("天选打工人 - %s", u.Name)
		} else {
			// 失败
			msg = fmt.Sprintf("%s 的%s(%s-%s)抢票失败了 ", u.Name, tomorrowStr, u.Line, u.Station)
			title = fmt.Sprintf("%s 抢票失败了", u.Name)
		}

		// 通知
		notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken, title, msg)
	}
}

func singleGetBalance(u *setting.User, date *time.Time) (balanceRes dto.MyBtiBalanceResponse, count int, err error) {
	// 循环n次抢票
	tomorrowStr := date.Format("20060102")
	for i := 1; i <= 5; i++ {
		_, err = http.NewRestyClient().R().
			SetResult(&balanceRes).
			SetBody(map[string]interface{}{
				"enterDate":          tomorrowStr,
				"lineName":           u.Line,
				"snapshotTimeSlot":   "0630-0930",
				"snapshotWeekOffset": 0,
				"stationName":        u.Station,
				"timeSlot":           u.Time,
			}).
			SetHeaders(map[string]string{"Authorization": u.AccessToken}).
			Execute("POST", "https://webapi.mybti.cn/Appointment/CreateAppointment")
		if err != nil {
			msg := fmt.Sprintf("%s抢的%s得票失败\n 错误信息: %s \n", u.Name, tomorrowStr, err.Error())
			log.Println(msg)

			notify.Pushplus(setting.GlobalConfig.Notify.PushplusToken, "抢票异常通知", msg)
			return balanceRes, i, err
		}
		if balanceRes.Balance > 0 {
			return balanceRes, i, err
		}
	}
	return balanceRes, -1, err
}
