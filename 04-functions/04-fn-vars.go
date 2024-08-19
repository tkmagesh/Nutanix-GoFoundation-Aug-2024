package main

import "fmt"

func main() {
	var greet func()
	fmt.Println("greet = ", greet)
	if greet == nil {
		fmt.Println("greet is not a func (but nil)")
	}
	greet = func() {
		fmt.Println("Hi there!")
	}
	greet()

	greet = func() {
		fmt.Println("Hello there!")
	}
	greet()
	/* ************* */

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello, %s!\n", userName)
	}
	sayHello("Magesh")
	/* ************** */
	var greetUser func(string, string) string
	greetUser = func(firstName, lastName string) string {
		return fmt.Sprintf("Hello %s %s, Have a nice day!\n", firstName, lastName)
	}
	fmt.Print(greetUser("Magesh", "Kuppan"))

	// use divide function as above
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
