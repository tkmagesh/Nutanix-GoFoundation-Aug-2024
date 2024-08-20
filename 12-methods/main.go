package main

import (
	"fmt"

	"github.com/tkmagesh/nutanix-gofoundation-aug-2024/12-methods/models"
)

func main() {

	/*
		var p models.Product = models.Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}


			var p2Ptr = &p
			p2Ptr.Name = "Fountain Pen"
			fmt.Printf("p.Name = %q, p2Ptr.Name = %q\n", p.Name, p2Ptr.Name)

			fmt.Println("After applying 10% discount")
			p.ApplyDiscount(10)
			fmt.Println(p.Format())
	*/
	// var newProduct = new(models.Product)
	// var newProduct = &models.Product{}
	var newProduct = models.NewProduct(101, "Pencil", 5)
	fmt.Println(newProduct)

	// var emp models.Employee
	var emp = new(models.Employee)
	fmt.Printf("%+v\n", emp)
}
