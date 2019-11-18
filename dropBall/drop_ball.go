package dropball

import (
	"fmt"
	"log"
	"time"
)

type ball struct {
	ID int
}

// Heaven 天堂
type Heaven struct {
	msgChan chan string
	personA *personA
	personB *personB
	personC *personC
}

type personA struct {
	handChan chan *ball
	heaven   *Heaven
}

type personB struct {
	handChan chan *ball
	heaven   *Heaven
}

type personC struct {
	handChan chan *ball
	heaven   *Heaven
}

// Lab 控制中心啟動
func (h *Heaven) Lab() {

	h.msgChan = make(chan string)

	h.personA = &personA{
		handChan: make(chan *ball, 5),
	}
	h.personB = &personB{
		handChan: make(chan *ball, 2),
	}
	h.personC = &personC{
		handChan: make(chan *ball, 2),
	}

	h.personA.heaven = h
	h.personB.heaven = h
	h.personC.heaven = h

	go h.personA.passBall()
	go h.personB.passBall()
	go h.personC.handleBall()
	go h.waitingCall()

	for i := 1; i <= 1000; i++ {

		ball := &ball{
			ID: i,
		}

		select {
		case h.personA.handChan <- ball:
			log.Println("A 拿到球 -> ", ball.ID)
		default:
			log.Println("A 沒手了")
		}

		time.Sleep(time.Second)
	}

}

func (h *Heaven) waitingCall() {
	for {
		select {
		case msg := <-h.msgChan:
			switch msg {
			case "B":
				log.Println("B call C")
			case "C":

				for i := 0; i < len(h.personC.handChan); i++ {
					fmt.Println("c Drop Ball -> ", <-h.personC.handChan)
					time.Sleep(time.Millisecond * 10)
				}

			}
		}
	}
}

func (a *personA) passBall() {

	for elme := range a.handChan {

		select {
		case a.heaven.personB.handChan <- elme:
			log.Println("B 拿到球 -> ", elme.ID)
		default:
			log.Println("B 沒手了")
			a.heaven.msgChan <- "B"
		}

	}

}

func (b *personB) passBall() {
	for elme := range b.handChan {
		select {
		case b.heaven.personC.handChan <- elme:
			log.Println("C 拿到球 -> ", elme.ID)

		default:
		Loop:
			for {
				select {
				case b.heaven.personC.handChan <- elme:
					log.Println("C 拿到球")
					break Loop
				default:
					log.Println("call C")
					b.heaven.msgChan <- "C"
					time.Sleep(time.Second / 2)
				}
			}
		}
	}
}

func (c *personC) handleBall() {
	for elme := range c.handChan {
		time.Sleep(time.Second * 2)
		log.Println("C 把球處理掉了 -> ", elme.ID)
	}
}
