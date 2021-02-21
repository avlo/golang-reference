package main

import (
	"fmt"
)

type adderfunc func(int, int) int

func nickAdderFunc() adderfunc {
	fmt.Println("111")
	return func(a int, b int) int {
		fmt.Println("2222")
		return a + b
	}
}

func functionAsMapValueMain() {
	// function type as map element
	addermap := map[int]adderfunc{}
	myfunc := nickAdderFunc() // or var myfunc = nickAdderFunc()
	addermap[0] = myfunc
	fmt.Printf("%d\n", addermap[0](1, 2))

	// function directly as map element
	addermap[1] = nickAdderFunc()
	fmt.Printf("%d\n", addermap[1](2, 2))
}

// functions as elememnt of struct

// BONUS BELOW
// function pointer as element of map
// function pointer as element of struct
