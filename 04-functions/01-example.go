package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	fmt.Print(getHello("Suresh"))
	fmt.Print(greet("Magesh", "Kuppan"))
	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Use only the quotient
	//  _ -> placeholder to receive a value that is not used
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/

	/*
		_, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, remainder = %d \n", r)
	*/

	// divide(100, 7)

}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayHello(userName string) {
	fmt.Printf("Hello, %s!\n", userName)
}

func getHello(userName string) string {
	return fmt.Sprintf("Hello, %s!\n", userName)
}

/*
func greet(firstName string, lastName string) string {
	return fmt.Sprintf("Hello %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) string {
	return fmt.Sprintf("Hello %s %s, Have a nice day!\n", firstName, lastName)
}

/*
func divide(multiplier, divisor int) (int, int) {
	quotient, remainder := multiplier/divisor, multiplier%divisor
	return quotient, remainder
}
*/

// named results (preferred for returning more than one result)
func divide(multiplier, divisor int) (quotient, remainder int) {
	// override the already declared & initialized variables
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
