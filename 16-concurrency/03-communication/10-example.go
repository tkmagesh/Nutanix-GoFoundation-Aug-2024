package main

import (
	"fmt"
	"sync"
)

func main() {
	// communication by sharing memory
	var result int
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, &result, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*result = x + y
}
