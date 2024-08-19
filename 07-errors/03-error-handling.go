package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	var divisor int
	fmt.Println("Enter the divisor")
	fmt.Scanln(&divisor)
	if q, r, err := divide(100, divisor); err != nil {
		fmt.Println("error occurred :", err)
		return
	} else {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, ErrDivideByZero
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

// named result
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient, remainder = x/y, x%y
	return
}
