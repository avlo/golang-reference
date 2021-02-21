package main

import (
	"fmt"
)

func pointers2(myStr *string) {
	fmt.Println(*myStr)
	*myStr = "new value"
	fmt.Println(*myStr)
}

func doSomething(i int) {
	fmt.Println(i)
}

type genfunc func(int)

func mapWithFuncs() {
	myMap := map[int]genfunc{}
	myMap[0] = doSomething
	myMap[0](0)

	var myfunk genfunc
	myfunk = doSomething
	myMap[1] = myfunk
	myMap[1](1)
}

func localmain2() {
	// val := "original"
	// pointers2(&val)
	// fmt.Println("notes2")
	mapWithFuncs()
}
