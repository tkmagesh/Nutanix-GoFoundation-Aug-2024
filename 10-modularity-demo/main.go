package main

import (
	"fmt"

	"github.com/tkmagesh/10-modularity-demo/calculator"
	/* "github.com/tkmagesh/10-modularity-demo/calculator/utils" */
	ut "github.com/tkmagesh/10-modularity-demo/calculator/utils"
)

func main() {
	fmt.Println("application executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op Count :", calculator.OpCount())
	fmt.Println("IsPrime(71) :", ut.IsPrime(71))
}
