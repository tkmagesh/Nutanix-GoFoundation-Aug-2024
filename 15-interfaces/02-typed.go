package main

import (
	"fmt"
	"math"
)

// ver 1.0

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

// ver 2.0
/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("argument has no Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("argument has no Area() method")
	}
}
*/

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// ver 3.0
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver 4.0

/*
func PrintShapeStats(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should ALSO be interface{ Perimeter() float64 }
}
*/

/*
func PrintShapeStats(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should ALSO be interface{ Perimeter() float64 }
}
*/

/*
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should ALSO be interface{ Perimeter() float64 }
}
*/

type ShapeStatsCalculator interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeStats(x ShapeStatsCalculator) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should ALSO be interface{ Perimeter() float64 }
}
func main() {
	c := Circle{17}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{12, 15}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

}
