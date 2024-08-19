package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no // data -> address
	fmt.Println("[main] &no = ", noPtr)

	// dereferencing (address -> data)
	fmt.Println(*noPtr)

	// in other words
	fmt.Println(no == *(&no))

	fmt.Println("[main] Before incrementing, no = ", no)
	// increment(no)
	incrementPtr(&no)
	fmt.Println("[main] After incrementing, no = ", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x int) {
	fmt.Println("[increment] &x = ", &x)
	x++
}

func incrementPtr(xPtr *int) {
	fmt.Println("[increment] xPtr = ", xPtr)
	(*xPtr)++
}

func swap(x, y *int) /* no return results */ {
	*x, *y = *y, *x
}
