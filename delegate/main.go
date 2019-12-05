package main

import (
	"goPra/delegate/printer"
	"goPra/delegate/receiver"
	"time"
)

func main() {

	p := &printer.Printer{}
	p.R = receiver.NewReceiver()

	p.PrintSomeThing()

	time.Sleep(time.Second * 1000)
}
