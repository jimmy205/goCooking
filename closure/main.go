package main

import (
	"fmt"
	"time"
)

func main() {

	n := getSN()

	go func() {
		t := time.NewTicker(time.Second)
		for {
			<-t.C
			n()
		}
	}()

	t := time.NewTicker(time.Millisecond * 300)
	for {
		<-t.C
		n()
	}
}

func getSN() func() int {
	n := 0
	return func() int {
		n++
		fmt.Println("n :", n)
		return n
	}
}
