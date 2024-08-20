package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... Hit ENTER!!\n", count)
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for idx := range count {
		wg.Add(1)
		go fn(wg, idx+1)
	}
	wg.Wait()
	fmt.Println("main completed!")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
