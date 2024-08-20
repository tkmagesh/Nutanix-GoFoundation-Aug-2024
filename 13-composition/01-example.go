package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	/* Dummy */
	Product
	/* Name   string */
	Expiry string
}

// overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

// PerishableProduct factory

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	// var milk PerishableProduct
	/*
		var milk = PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Milk",
				Cost: 48,
			},
			Expiry: "1 Day",
		}
	*/

	// create an instance using the factory function
	var milk = NewPerishableProduct(100, "Milk", 48, "1 Day")
	fmt.Printf("%#v\n", milk)

	// fmt.Println(milk.Product.Name)
	// fmt.Println(milk.Name)
	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
