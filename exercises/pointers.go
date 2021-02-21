package main

import "fmt"

func ref(ptr *int) {
	*ptr++
}

func pointersmain() {
	i := 1
	fmt.Println(i)

	ref(&i)
	fmt.Println(i)

	var j *int
	j = &i
	fmt.Println(*j)
	ref(j)
	fmt.Println(*j)
}
