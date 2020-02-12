package main

import (
	"goPra/corvus/constant"
	"goPra/corvus/monitor"
	"goPra/corvus/package/nsq"

	"time"
)

func main() {

	// load.LoadConfig()

	// 開始監視
	monitoring()

	// 定期回傳訊息到背景確認corvus沒死掉
	nsqProducer := nsq.NewProducer()

	t := time.NewTicker(time.Minute * 5)
	for {
		nsqProducer.SendToBackground(constant.NsqTopicCorvusMonitor, []byte("corvus"))
		<-t.C
	}
}

func monitoring() {
	// 取得監視器
	m := monitor.NewMonitor()

	// 監控服務是否正常
	nsq.SetReceiverNSQ(constant.NsqTopicServerMonitor, constant.NsqChannelServerName, m.Monitor())
	go m.CheckAlive()
}
