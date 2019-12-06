package main

import "fmt"

var cmMap = map[string]func(){
	"A": printA,
	"B": printB,
	"C": printC,
	"D": printD,
	"E": printE,
	"F": printF,
}

func main() {
	for k := range cmMap {
		fn := cmMap[k]
		fn()
	}
}

func printA() {
	fmt.Println("A")
}

func printB() {
	fmt.Println("B")
}

func printC() {
	fmt.Println("C")
}

func printD() {
	fmt.Println("D")
}

func printE() {
	fmt.Println("E")
}

func printF() {
	fmt.Println("F")
}
