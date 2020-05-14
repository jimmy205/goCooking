package main

import (
	"goPra/lesson/hotel"

	"math/rand"
	"strconv"
	"time"
)

func main() {
	hotel.InitManager()

	id := 0

	go func() {
		for {
			id++
			torist := hotel.NewTourist("normal_"+strconv.Itoa(id), true)
			hotel.DefaultManager.JoinRoom(torist)

			// 隨機的時間
			randTime := rand.New(rand.NewSource(time.Now().Unix())).Intn(10)
			if randTime <= 0 {
				randTime = 10
			}

			time.Sleep(time.Second * time.Duration(randTime))
		}
	}()

	t := time.NewTicker(time.Second * 10)
	for {
		<-t.C
		hotel.DefaultManager.ShowRoomDetail()
	}

}
