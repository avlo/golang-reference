package main

import "fmt"

type mystruct struct {
	myint int
}

func (m *mystruct) getIntRef() int {
	return m.myint
}

func (m mystruct) getIntVal() int {
	return m.myint
}

func (m *mystruct) setIntRef(i int) {
	m.myint = i
}

func (m mystruct) setIntVal(i int) {
	m.myint = i
}

func functionsmain() {
	nickstruct := mystruct{
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
