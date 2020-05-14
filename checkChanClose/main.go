package main

import "log"

func main() {

	a := make(chan int, 2)

	a <- 1

	// close(a)
	// ok := sendToChan(a, 2)

	select {
	case ok := <-sendToChanForSel(a, 3):
		log.Println("1 -> ", ok)
	}
}

func sendToChan(ch chan<- int, value int) bool {
	defer func() {
		recover()
	}()

	ch <- value

	return true
}

func sendToChanForSel(ch chan<- int, value int) chan bool {
	ok := make(chan bool)

	go func() {
		defer close(ok)
		defer func() {
			recover()
			log.Println("2")
			ok <- false
		}()

		log.Println("3")
		ch <- value
		ok <- true
	}()

	return ok
}
