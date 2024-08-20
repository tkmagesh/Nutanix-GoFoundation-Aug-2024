package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
}

func genData(ch chan int) {
	count := rand.Intn(20)
	fmt.Println("Count :", count)
	for i := range count {
		ch <- (i + 1) * 10
	}
	close(ch)
}
