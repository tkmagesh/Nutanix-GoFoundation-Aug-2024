package calculator

import "github.com/fatih/color"

// private
var opCount map[string]int

func init() {
	color.Red("calculator initialized - [calc.go]")
	// fmt.Println("calculator initialized - [calc.go]")
	opCount = make(map[string]int)
}

// public
func OpCount() map[string]int {
	return opCount
}
