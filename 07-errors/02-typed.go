package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrRandom error = errors.New("random error")

func main() {
	err := fn()
	if err == ErrRandom {
		fmt.Println("random error occurred")
		return
	}
	if err != nil {
		fmt.Printf("error executing fn, err = %v\n", err)
		return
	}
}

func fn() error {
	randomNo := rand.Intn(20)
	if randomNo%2 == 0 {
		return ErrRandom
	}
	fmt.Println("[fn] executed without errors")
	return nil
}
