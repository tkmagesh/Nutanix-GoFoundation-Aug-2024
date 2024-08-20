package calculator

import "fmt"

func init() {
	fmt.Println("calculator initialized - [add.go]")
}

func Add(x, y int) int {
	opCount["add"]++ // declared in 'calc.go'
	return x + y
}
