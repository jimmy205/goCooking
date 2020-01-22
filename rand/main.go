package main

import (
	realRand "crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	for {
		b := big.NewInt(10000)
		rr, err := realRand.Int(realRand.Reader, b)
		fmt.Println("rr :", rr.Uint64(), err)
	}

}
