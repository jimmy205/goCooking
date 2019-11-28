package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const addWorker = 500000
const cutWorker = 850000
const perProduce = "600000"

var customer = 0

type random struct {
	randoms       chan int
	workerList    []chan bool
	workerListCtx []context.CancelFunc
	maxRandoms    int
	maxWorker     int
	minWorker     int
	useRandom     int
	produceRandom int
}

func main() {

	r := &random{
		randoms:       make(chan int, 3000000),
		workerList:    []chan bool{},
		maxRandoms:    3000000,
		maxWorker:     5,
		minWorker:     5,
		useRandom:     0,
		produceRandom: 0,
	}

	// test get random time
	// st := time.Now()
	// for worker := 0; worker < 10; i++ {
	// 	go r.testfunc()
	// }
	// fmt.Scanln()
	// fmt.Println(" ====== total time ====== : ", time.Since(st))

	go r.jobCenter()

	go r.jobCenterCtx()

	go r.factorySpeaker()

	go inputReader()

	t := time.NewTicker(time.Millisecond * 5)
	for {
		<-t.C
		for i := 0; i < customer; i++ {
			go r.randomShop()
		}
	}

}

func (r *random) testfunc() {

	t := time.NewTicker(time.Millisecond * 1)

	for {
		<-t.C
		st := time.Now()

		randoms := randomResource()
		fmt.Println("slice -> ", len(randoms), "excute time : ", time.Now().Sub(st))
	}

}

func inputReader() {
	t2 := time.NewTicker(time.Second * 2)

	for {
		<-t2.C
		fmt.Scan(&customer)
	}
}

func (r *random) jobCenterCtx() {

	for {
		// 不知道為甚麼要睡一下，不睡會壞掉
		time.Sleep(time.Millisecond * 1)

		randomsLen := len(r.randoms)
		workerLen := len(r.workerListCtx)

		if randomsLen < addWorker && workerLen < r.maxWorker {
			ctx, cancel := context.WithCancel(context.Background())
			r.workerListCtx = append(r.workerListCtx, cancel)
			go r.randomFactory(ctx, nil)
		}

		if randomsLen > cutWorker && workerLen > r.minWorker {
			cancel := r.workerListCtx[0]
			cancel()
			r.workerListCtx = append(r.workerListCtx[:0], r.workerListCtx[1:]...)
		}

	}
}

func (r *random) jobCenter() {

	for {
		// 不知道為甚麼要睡一下，不睡會壞掉
		time.Sleep(time.Millisecond * 10)

		randomsLen := len(r.randoms)
		workerLen := len(r.workerList)

		if randomsLen < addWorker && workerLen < r.maxWorker {
			workSignal := make(chan bool)
			r.workerList = append(r.workerList, workSignal)
			go r.randomFactory(nil, workSignal)
		}

		if randomsLen > cutWorker && workerLen > r.minWorker {
			r.workerList[0] <- true
			close(r.workerList[0])
			r.workerList = append(r.workerList[:0], r.workerList[1:]...)
		}
	}
}

func (r *random) randomFactory(ctx context.Context, working chan bool) {

	if ctx != nil {
		for {

			select {
			case <-ctx.Done():
				return
			default:
				r.packingRandoms()
			}

		}

	} else {
	Loop:
		for {

			select {
			case _, ok := <-working:
				if ok {
					break Loop
				} else {
					log.Println("worker channel closed already")
					break Loop
				}
			default:
				r.packingRandoms()
			}

		}
	}

}

func (r *random) packingRandoms() {

	numSlice := randomResource()

	for i := range numSlice {
		if len(r.randoms) == r.maxRandoms {
			break
		}
		r.randoms <- numSlice[i]
		r.produceRandom++
	}

}

func (r *random) randomShop() {
	<-r.randoms
	r.useRandom--
}

func (r *random) factorySpeaker() {

	t := time.NewTicker(time.Second)
	t2 := time.NewTicker(time.Second * 10)
	lastProduce := 0
	lastUse := 0
	lastCapacity := 0
	capacity := 0
	count := 0
	for {

		select {
		case <-t.C:
			log.Println("now randoms     :", len(r.randoms))
			log.Println("now customer    :", customer)
			log.Println("we have ctx     :", len(r.workerListCtx), " workers")
			log.Println("we have channel :", len(r.workerList), " workers")
			// log.Println("we randoms produce count      :", r.produceRandom)
			log.Println("produce this time :", r.produceRandom-lastProduce)
			log.Println("use random this time :", r.useRandom-lastUse)
			capacity = (r.produceRandom - lastProduce) + (r.useRandom - lastUse)
			log.Println("capacity : ", capacity)
			fmt.Println("----- ----- ---- ----- ----- ----")
			if capacity < lastCapacity && capacity < 10000 {
				count++
				if count > 10 {
					log.Println("its over charging!!!")
				}
			}
			lastCapacity = capacity
			lastProduce = r.produceRandom
			lastUse = r.useRandom
		case <-t2.C:
			log.Println("10 sec over : ", count)
			count = 0
		}

	}
}

func randomResource() []int {
	var out bytes.Buffer
	cmd := exec.Command("head", "-c", "20", "/dev/hwrng")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

		var isSuccess bool

		for !isSuccess {
			//本機mac開發無/dev/hwrng
			cmd := exec.Command("head", "-c", perProduce, "/dev/urandom")
			cmd.Stdout = &out
			err2 := cmd.Run()
			if err2 != nil {
				log.Println("error", "error_test: "+err.Error())
				continue
			}
			isSuccess = true
		}
	}

	n := bytes.Split(out.Bytes(), []byte(" "))

	var numSlice []int
	for _, v := range n {
		if len(v) > 4 {
			randNum := binary.BigEndian.Uint32(v)
			numSlice = append(numSlice, int(randNum))
		}
	}

	return numSlice
}
