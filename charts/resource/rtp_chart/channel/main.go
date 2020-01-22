package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bufAndNotBuf()
}

func bufAndNotBuf() {

	/*
	 * 如果有給長度，會先塞值進去等人取
	 * (every 1 sec) success  ->  get success , success  ->  get success
	 */
	ch := make(chan bool, 1)

	/*
	 * 如果用預設值，會等到有人要取的時候才會塞值
	 * (after 10 sec) get success -> success , get success -> success
	 */
	// ch2 := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- true
			log.Println("success")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 10)
			<-ch
			log.Println("get success")
		}
	}()

	fmt.Scanln()
}
