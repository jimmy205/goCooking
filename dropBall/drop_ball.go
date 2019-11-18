package dropball

import (
	"fmt"
	"time"
)

type ball struct {
	to string
}

// Heaven 天堂
type Heaven struct {
	msgChan chan string
	personA *person
	personB *person
	personC *person
}

type person struct {
	handChan chan *ball
}

// Lab 控制中心啟動
func (h *Heaven) Lab() {

	h.msgChan = make(chan string)

	h.personA = &person{
		handChan: make(chan *ball, 5),
	}
	h.personB = &person{
		handChan: make(chan *ball, 2),
	}
	h.personC = &person{
		handChan: make(chan *ball, 2),
	}

	go h.getBall()

	for {
		time.Sleep(time.Second)

		ball := &ball{
			to: "B",
		}

		select {
		case h.personA.handChan <- ball:
			fmt.Println("A got ball,to B")
		default:
			fmt.Println("A miss ball ... ?")
		}

	}

}

func (h *Heaven) getBall() {
	for {

		msg := &ball{}
		select {
		case msg = <-h.personA.handChan:
		case msg = <-h.personB.handChan:
			msg.to = "C"
		case msg = <-h.personC.handChan:
			fmt.Println("c got Ball,handle ball")
			for i := 0; i < 3; i++ {
				fmt.Println("handling -> ", i)
				time.Sleep(time.Second)
			}
		}

		h.handleBall(msg)
	}
}

func (h *Heaven) handleBall(msg *ball) {

	switch msg.to {
	case "B":
		select {
		case h.personB.handChan <- msg:
			fmt.Println("B got ball,to C")
		default:
			fmt.Println("B miss ball,A call god")
		}
	case "C":
		select {
		case h.personC.handChan <- msg:
			fmt.Println("C got ball,Ready Handle it")

		default:
			fmt.Println("C miss ball,B call god")
		}

	}

}
