package main

import (
	"log"

	"golang.org/x/sync/singleflight"
)

var reqSingleFlight singleflight.Group

func main() {
	sfDo()
}

func sfDo() {
	reqSingleFlight.Do("123", func() (interface{}, error) {
		return dododo()
	})

}

func dododo() (string, error) {
	log.Println("iiiii")
	return "222", nil
}
