package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
}
