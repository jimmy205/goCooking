package receiver

import (
	"fmt"
	"time"
)

import printSpec "gopra/delegate/printspec"

// Receiver 接收東西
type Receiver struct {
	s            string
	rHandlePrint func(string)
}

// NewReceiver 新的接收器
func NewReceiver() printSpec.IPrintCore {
	r := &Receiver{
		s:            "a",
		rHandlePrint: func(string) {},
	}

	go r.ReceiveSomething()
	go r.PrintSomething()
	return r
}

// ReceiveSomething 接收東西
func (r *Receiver) ReceiveSomething() {
	// r.rHandlePrint(s)
	for {
		fmt.Scanln(&r.s)
	}
}

// PrintSomething 印東西
func (r *Receiver) PrintSomething() {
	t := time.NewTicker(time.Second)
	for {
		<-t.C
		if r.s != "a" {
			r.rHandlePrint(r.s)
		}
	}
}
