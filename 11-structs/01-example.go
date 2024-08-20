package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var p Product
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/

	// var p Product = Product{100, "Pen", 10}
	// var p Product = Product{Id: 100, Name: "Pen", Cost: 10}
	// var p Product = Product{Id: 100, Name: "Pen"}
	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	/*
		var p2 = p //create a copy
		p2.Name = "Fountain Pen"
		fmt.Printf("p.Name = %q, p2.Name = %q\n", p.Name, p2.Name)
	*/

	var p2Ptr = &p
	/*
		(*p2Ptr).Name = "Fountain Pen"
		fmt.Printf("p.Name = %q, p2Ptr.Name = %q\n", p.Name, (*p2Ptr).Name)
	*/

	p2Ptr.Name = "Fountain Pen"
	fmt.Printf("p.Name = %q, p2Ptr.Name = %q\n", p.Name, p2Ptr.Name)

	fmt.Println("After applying 10% discount")
	ApplyDiscount(&p, 10)
	// fmt.Printf("%+v\n", p)
	fmt.Println(Format(p))
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
