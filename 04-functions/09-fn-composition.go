package main

import (
	"fmt"
	"log"
	"time"
)

type OperationFn func(int, int)

func main() {
	/*
		profiledAdd := getProfileOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfileOperation(subtract)
		profiledSubtract(100, 200)
	*/

	// logging & profiling

	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfileOperation(logAdd)
		profiledLogAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		profiledLogSubtract := getProfileOperation(logSubtract)
		profiledLogSubtract(100, 200)

		getProfileOperation(getLogOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		}))(100, 200)
	*/
	// re-wiring
	/*
		add := getLogOperation(add)
		add = getProfileOperation(add)
		add(100, 200)
	*/

	add := compose(add, getLogOperation, getProfileOperation)
	add(100, 200)

	subtract := compose(subtract, getLogOperation, getProfileOperation)
	subtract(100, 200)
}

func compose(op OperationFn, wrappers ...func(OperationFn) OperationFn) OperationFn {
	for idx := len(wrappers) - 1; idx >= 0; idx-- {
		wrapper := wrappers[idx]
		op = wrapper(op)
	}
	return op
}

/*
func getProfileOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
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
*/

// typed
func getProfileOperation(op OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func getLogOperation(op OperationFn) OperationFn {
	var opWrapper OperationFn
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
