package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("panic occurred : err = ", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	for {
		var divisor int
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		q, r, err := divideWrapper(100, divisor)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try again!")
			continue
		}
		fmt.Printf("[main] - Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// wrapper to conver the panic from "divide()" into an error using the deferred function
func divideWrapper(x, y int) (q, r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	q, r = divide(x, y)
	return
}

// 3rd party api
func divide(multiplier, divisor int) (quotient, remainder int) {
	defer func() {
		fmt.Println("	[divide] - deferred")
	}()
	fmt.Println("[divide] - calculating quotient")
	if divisor == 0 {
		panic(ErrDivideByZero)
	}
	quotient = multiplier / divisor
	fmt.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	fmt.Println("[divide] - return the result")
	return
}
