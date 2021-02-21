package main

import "fmt"

type myRefStruct struct {
	myint int
}

func (m *myRefStruct) getIntRef() int {
	return m.myint
}

func (m myRefStruct) getIntVal() int {
	return m.myint
}

func (m *myRefStruct) setIntRef(i int) {
	m.myint = i
}

func (m myRefStruct) setIntVal(i int) {
	m.myint = i
}

func methodsmain() {
	nickstruct := myRefStruct{
		myint: 1,
	}

	fmt.Println(nickstruct.myint) // should be 1
	nickstruct.setIntVal(1)
	fmt.Println(nickstruct.getIntVal()) // should be 1
	fmt.Println(nickstruct.myint)       // should be 1
	fmt.Println()

	nickstruct.setIntRef(2)
	fmt.Println(nickstruct.getIntRef()) // should be 2
	fmt.Println(nickstruct.myint)       // should be 2

}
