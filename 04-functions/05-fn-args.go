package main

import "fmt"

func main() {
	// decide which function to be executed using the "exec" function
	/*
		exec("F1")
		exec("f2")
		exec("nonExistentFunc")
	*/

	execFn(f1)
	execFn(f2)
	// execFn(nonExistentFunc)
	execFn(f3)
	execFn(func() {
		fmt.Println("anonymous func invoked")
	})
}

func exec(fnName string) {
	// executing a function
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("Invalid argument")
	}
}

func execFn(fn func()) {
	fn()
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
