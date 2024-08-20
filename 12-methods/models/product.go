package models

type Product struct {
	Id   int
	Name string
	Cost float64
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}
