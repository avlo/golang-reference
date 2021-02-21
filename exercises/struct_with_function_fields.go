package main

import "fmt"

type structAdder func(int, int) int

type myStruct struct {
	structFunction structAdder
}

func callAdderClusure() structAdder {
	return func(i int, j int) int {
		return i + j
	}
}

func structWithFunctionFieldsMain() {

	// functions as elememnt of struct
	localStruct := myStruct{}
	localStruct.structFunction = callAdderClusure()
	fmt.Printf("%v", localStruct.structFunction(1, 2))

	// BONUS BELOW
	// function pointer as element of map
	// function pointer as element of struct
}
