package main

import "fmt"

// Geometry comment
type Geometry interface {
	Area(base int, height int) int
}

type square struct{}
type triangle struct{}
type circle struct{}

func (s square) Area(base int, height int) int {
	return base * height
}

func (t triangle) Area(base int, height int) int {
	return 1 / 2 * (base * height)
}

func calcGeometry(g Geometry) int {
	return g.Area(2, 2)
}

func interfacesmain() {
	s := square{}
	t := triangle{}
	c := circle{}
	calcGeometry(c)
	fmt.Println(calcGeometry(s))
	fmt.Println(calcGeometry(t))
}
