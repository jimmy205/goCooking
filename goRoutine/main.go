package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// chapter1() // channel ok meaning
	// chapter2() // random pick

	// fmt.Scanln()
}

func chapter1() {
	c := make(chan int, 2)

	go func() {
		c <- 1
		c <- 1
	}()

	go func() {
		var i int
		for {
			time.Sleep(time.Second)
			i++

			if i == 2 {
				close(c)
			}
			n, ok := <-c
			log.Println("n -> ", n, "ok -> ", ok)
		}
	}()

	fmt.Scanln()
}

func chapter2() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan string, 10)

	go func() {
		for i := 0; i < 100; i++ {
			ch2 <- "ch2"
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ch3 <- "ch3"
		}
	}()

	go func() {
		for {

			// random pick
			select {
			case s, ok := <-ch1:
				if ok {
					log.Println("s -> ", s)
				}
			case s, ok := <-ch2:
				if ok {
					log.Println("s -> ", s)
				}
			case s, ok := <-ch3:
				if ok {
					log.Println("s -> ", s)
				}
			}
		}

	}()

	fmt.Scanln()
}
