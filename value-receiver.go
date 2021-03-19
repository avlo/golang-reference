package main

import (
	"fmt"
)

type greet struct {
	greet string
}

// receiver value type (like regular value parameter) doesn't update externally
func (g greet) HelloVal() {
	g.greet = "value"
}

// setters in golang always need to be ref type!
// receiver pointer/ref type (like pointer/ref parameter) DOES update externally
func (g *greet) HelloPtr() {
	g.greet = "reference"
}

func main() {

	greeting := greet{"greek"}
	fmt.Println(greeting.greet)
	greeting.HelloVal()
	fmt.Println(greeting.greet)

	greeting.HelloPtr()
	fmt.Println(greeting.greet)

	(&greeting).HelloPtr()
	fmt.Println(greeting.greet)
}

