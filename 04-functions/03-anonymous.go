/*
anonymous functions
  - functions with no name
  - created & invoked immediately
*/
package main

import "fmt"

func main() {

	// sayHi as anonymous func
	func() {
		fmt.Println("Hi there!")
	}()

	/* 1 parameter, 0 return values */
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(addResult)
}
