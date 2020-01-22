package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

func main() {
	b := big.NewInt(100)
	r, err := rand.Int(rand.Reader, b)
	log.Println("r :", r.Uint64(), "err :", err)
}
