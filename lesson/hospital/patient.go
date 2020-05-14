package hospital

import (
	"math/rand"
	"time"
)

// Patient 病患
type Patient struct {
	name           string
	sendToHospital bool
	observeDays    int // 觀察天數
}

// NewPatient 新的病患
func NewPatient(name string) *Patient {

	observeDays := rand.New(rand.NewSource(time.Now().Unix())).Intn(14)

	p := &Patient{
		name:        name,
		observeDays: observeDays,
	}

	return p
}
