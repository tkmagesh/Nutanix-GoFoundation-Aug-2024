package main

import (
	"fmt"
	"sync"
)

func main() {
	// share memory by communicating
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, ch, wg)
	result := <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- x + y
}
