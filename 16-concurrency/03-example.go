package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 20 {
		wg.Add(1) // incrementing the wg counter by 1
		go f1(wg)
	}
	f2()
	wg.Wait() // block the execution of "main()" until the wg counter becomes "0" (default)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
