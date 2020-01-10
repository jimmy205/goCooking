package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover get!!!")
		}
	}()

	time.Sleep(time.Second * 3)
	panic("")
}
