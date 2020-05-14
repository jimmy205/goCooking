package hospital

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Hospital 醫院
type Hospital struct {
	patientNum            int
	patientMap            map[*Patient]bool
	patientMapLock        *sync.RWMutex
	sendPatientHandler    func([]string)
	sendPatientTestResult func(map[string]bool)
}

const id = iota

// NewHospital 新醫院
func NewHospital() IHostpital {
	h := &Hospital{
		patientMap:            make(map[*Patient]bool),
		sendPatientHandler:    func([]string) {},
		sendPatientTestResult: func(map[string]bool) {},
		patientMapLock:        new(sync.RWMutex),
	}

	go h.running()

	return h
}

func (h *Hospital) running() {
	sendT := time.NewTicker(time.Second * 5)
	checkT := time.NewTicker(time.Second * 2)

	for {

		select {
		case <-sendT.C:
			h.sendPatientHandler(h.NewFoundPatient())
		case <-checkT.C:

		}

	}
}

// NewFoundPatient 發現新的病患
func (h *Hospital) NewFoundPatient() (namelist []string) {
	// 產生亂數
	randNum := rand.New(rand.NewSource(time.Now().Unix())).Intn(5)

	// 名單
	for i := 0; i <= randNum; i++ {
		h.patientNum++
		p := NewPatient("patient_" + strconv.Itoa(h.patientNum))
		namelist = append(namelist, p.name)
		h.patientMap[p] = true
	}

	return
}

// CheckResult 檢查結果
func (h *Hospital) CheckResult() {
	h.patientMapLock.Lock()
	for p := range h.patientMap {

		p.observeDays--
		if p.observeDays <= 0 {
			h.sendPatientTestResult(map[string]bool{
				p.name: false,
			})
		}

		// 取亂數看有沒有病
		randNum := rand.New(rand.NewSource(time.Now().Unix())).Intn(100)
		if randNum%2 == 0 {
			p.observeDays += 3
			h.sendPatientTestResult(map[string]bool{
				p.name: true,
			})
		}
	}
	h.patientMapLock.Unlock()
}
