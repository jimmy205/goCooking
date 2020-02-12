package main

import (
	"log"
	"reflect"
	"time"
)

// Service 監視器
type Service struct {
	// Pelican    detail
	// Background detail
	// Raven      detail
	Servers []detail
}

type detail struct {
	ServiceName string
	UpdateTime  time.Time
}

func main() {

	setStruct()
	return
	// st := &Service{
	// 	Pelican: detail{
	// 		ServiceName: "pelican",
	// 		UpdateTime:  time.Now(),
	// 	},
	// 	Background: detail{
	// 		ServiceName: "background",
	// 		UpdateTime:  time.Now(),
	// 	},
	// 	Raven: detail{
	// 		ServiceName: "raven",
	// 		UpdateTime:  time.Now(),
	// 	},
	// }

	// v := reflect.ValueOf(st).Elem()
	// for i := 0; i < v.NumField(); i++ {

	// 	log.Println("v :", v.Field(i).FieldByName("UpdateTime").Interface().(time.Time))
	// }

}

func setStruct() {

	s := &Service{}
	r := reflect.ValueOf(s).Elem()
	f := reflect.Indirect(r).FieldByName("Servers")

	st := []string{
		"raven",
		"pelican",
	}

	ss := []detail{}
	for _, s := range st {
		ss = append(ss, detail{
			ServiceName: s,
			UpdateTime:  time.Now(),
		})
	}
	f.Set(reflect.ValueOf(ss))

	log.Println("s :", s)
}
