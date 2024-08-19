package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// execute the function
	if fn := getFn(); fn != nil {
		fn()
		return
	}
	fmt.Println("Not a function to invoke")
}

func getFn() func() {
	// deciding the function to execute
	var randomNo = rand.Intn(20)
	switch {
	case randomNo%2 == 0:
		return f1
	case randomNo%3 == 0:
		return f2
	case randomNo%5 == 0:
		return func() {
			fmt.Println("anonymous func invoked")
		}
	default:
		return nil
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3() {
	fmt.Println("f3 invoked")
}
