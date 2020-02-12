package telegram

import (
	"os"
	"time"
)

// SendNormalMsg 傳送普通訊息
func SendNormalMsg(msg string) {
	send(msg, false)
}

// SendNoticeMsg 傳送注意的訊息
func SendNoticeMsg(serviceName string, lastTime time.Time) {

	s := "『 ⚠️⚠️⚠️ *** Cuckoo提醒 *** ⚠️⚠️⚠️ 』\n "
	s += "〔 *** 環境 *** 〕: \n ``` " + os.Getenv("PROJECT_ENV") + "``` \n"
	s += "〔 *** 沒有收到的服務名稱 *** 〕: \n ``` " + serviceName + " ``` \n"
	s += "〔 *** 最後一次收到訊息的時間 *** 〕: \n ``` " + lastTime.Format("2006-01-02 15:04:05") + " ``` \n"

	// 發送訊息
	send(s, true)
}
