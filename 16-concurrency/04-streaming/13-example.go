package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan []int)
	go genData(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func genData(ch chan []int) {
	ch <- []int{10}
	time.Sleep(500 * time.Millisecond)
	ch <- []int{10, 20}
	time.Sleep(500 * time.Millisecond)
	ch <- []int{10, 20, 30, 40}
	time.Sleep(500 * time.Millisecond)
	ch <- []int{10}
	time.Sleep(500 * time.Millisecond)
	ch <- []int{10}
	time.Sleep(500 * time.Millisecond)
}
