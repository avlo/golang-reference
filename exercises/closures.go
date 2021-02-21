package main

import (
	"fmt"
)

// feels like a java interface
type mathoperation func(int, int) int

func add() mathoperation {
	c := 10
	return func(a int, b int) int {
		return a + b + c
	}
}

func sub() mathoperation {
	c := 10
	return func(a int, b int) int {
		return c - a - b
	}
}

func addWithParameter(d int) mathoperation {
	c := 10
	return func(a int, b int) int {
		return a + b + c + d
	}
}

func subWithParameter(d int) mathoperation {
	c := 10
	return func(a int, b int) int {
		return d - c - a - b
	}
}

type adder func(int, int) int

func localmain() {
	// closures
	adder := add() // or var adder add()
	fmt.Println(adder(2, 2))
	subtractor := sub() // or var subtractor sub()
	fmt.Println(subtractor(2, 2))

	fmt.Println(addWithParameter(100)(2, 2))
	fmt.Println(subWithParameter(100)(2, 2))
}

// functions as elememnt of struct

// BONUS BELOW
// function pointer as element of map
// function pointer as element of struct
