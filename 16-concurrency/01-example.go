package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of "f1" through the "Go Scheduler"
	f2()
	// blocking the execution of "main()" so that the scheduler can look for other goroutines scheduled and execute them

	// DO NOT DO THE FOLLOWING (Poor man's synchronization techniques) in production
	time.Sleep(4 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
