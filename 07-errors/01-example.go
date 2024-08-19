package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	err := fn()
	if err != nil {
		fmt.Printf("error executing fn, err = %v\n", err)
		return
	}
}

func fn() error {
	randomNo := rand.Intn(20)
	if randomNo%2 == 0 {
		return errors.New("random error")
	}
	fmt.Println("[fn] executed without errors")
	return nil
}
