package dropball

import (
	"time"
)

// Heaven 天堂
type Heaven struct {
	msgChan chan string
	personA *person
	personB *person
	personC *person
}

type person struct {
	handChan chan string
}

// Lab 控制中心啟動
func (h *Heaven) Lab() {

	go h.peopleWaiting()

	h.msgChan = make(chan string)

	h.personA = &person{
		handChan: make(chan string),
	}
	h.personB = &person{
		handChan: make(chan string, 2),
	}
	h.personC = &person{
		handChan: make(chan string, 2),
	}

	ballDropping := time.NewTicker(time.Second)
	for {

		select {
		// 發球機
		case <-ballDropping.C:
			h.personA.handChan <- "ball"
		case role := <-h.msgChan:

			switch role {
			case "b":
				if len(h.personC.handChan) == 2 {
					go func() {
						h.callLab("c")
					}()
				}
			case "c":
				for {
					<-h.personC.handChan
				}
			}
		}

	}

}

// peopleWaiting 人在等待球
func (h *Heaven) peopleWaiting() {

	// 如果A要接球，可是B還沒接，打電話給實驗室，叫實驗室跟B說快點接球
	// C如果被摧促，可以把球丟掉

	for {

		select {
		case b := <-h.personA.handChan:
			h.personB.handChan <- b

			// 如果B在忙沒辦法接球，打給實驗室
			if len(h.personB.handChan) == 2 {
				go func() {
					h.callLab("b")
				}()
			}
		case b := <-h.personB.handChan:

			h.personC.handChan <- b
		case <-h.personC.handChan:
			time.Sleep(time.Second * 2)

		}
	}
}

// callLab 呼叫上帝
func (h *Heaven) callLab(role string) {
	h.msgChan <- role
}
