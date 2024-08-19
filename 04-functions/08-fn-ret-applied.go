package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)
		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// rewiring
	add := getLogOperation(add)
	subtract := getLogOperation(subtract)
	add(100, 200)
	subtract(100, 200)

}

func getLogOperation(op func(int, int)) func(int, int) {
	var opWrapper func(int, int)
	opWrapper = func(x, y int) {
		log.Println("Operation started...")
		op(x, y)
		log.Println("Operation completed...!")
	}
	return opWrapper
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
