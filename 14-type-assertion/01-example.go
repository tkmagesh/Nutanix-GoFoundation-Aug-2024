package main

import "fmt"

type Product struct {
}

func main() {
	var x interface{}
	x = 100
	x = "Ut enim excepteur cupidatat ad in."
	x = 99.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	x = "100"
	// y := x + 200
	// y := x.(int) + 200
	// fmt.Println(y)

	// type assertion
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type assertion using "type switch"
	// x = 100
	// x = "Ut enim excepteur cupidatat ad in."
	// x = 99.99
	// x = true
	// x = Product{}
	x = struct{}{}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, x * 0.99 =", val*0.99)
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	case Product:
		fmt.Println("x is a product")
	default:
		fmt.Println("x is of unknown type")
	}

}
