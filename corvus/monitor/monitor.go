package monitor

import (
	"log"
	"time"

	FishTelegram "goPra/corvus/package/telegram"
)

// Service 監視器
type Service struct {
	test       detail
	Pelican    detail
	Background detail
}

type detail struct {
	serviceName string
	updateTime  time.Time
}

// NewMonitor 新的監視器
func NewMonitor() *Service {
	return &Service{
		test: detail{
			serviceName: "test",
			updateTime:  time.Now(),
		},
		Pelican: detail{
			serviceName: "pelican",
			updateTime:  time.Now(),
		},
		Background: detail{
			serviceName: "background",
			updateTime:  time.Now(),
		},
	}
}

// Monitor 測試監視器
func (s *Service) Monitor() func(msg []byte) {
	return func(msg []byte) {

		sn := string(msg)
		switch sn {
		case "pelican":
			s.Pelican.updateTime = time.Now()
		case "background":
			log.Println("Get background")
			s.Background.updateTime = time.Now()
		default:
		}

	}
}

// CheckAlive 確認還活著
func (s *Service) CheckAlive() {

	t := time.NewTicker(time.Minute * 5)
	for {
		<-t.C
		if time.Now().Sub(s.Pelican.updateTime).Minutes() > 10 {
			FishTelegram.SendNoticeMsg(s.Pelican.serviceName, s.Pelican.updateTime)
		}

		if time.Now().Sub(s.Background.updateTime).Minutes() > 10 {
			FishTelegram.SendNoticeMsg(s.Background.serviceName, s.Background.updateTime)
		}
	}

}

// ShowService 顯示目前的資訊
func (s *Service) ShowService() {
	log.Println("name :", s.Pelican.serviceName, " , update time :", s.Pelican.updateTime)
	log.Println("name :", s.Background.serviceName, " , update time :", s.Background.updateTime)
}
