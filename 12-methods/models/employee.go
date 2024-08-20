package models

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

type Organization struct {
	Id   int
	Name string
	City string
}
