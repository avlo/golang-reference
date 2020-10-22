package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {
		fmt.Println("in thread 2: 5sec delty to channel write")
		time.Sleep(time.Second * 5)
		ch <- 2
	}(ch)

	go func(ch chan int) {
		fmt.Println("in thread 1: 2sec delay to channel write")
		time.Sleep(time.Second * 2)
		ch <- 1
	}(ch2)

	fmt.Println("in main: blocking from 5 sec thread")
	fmt.Println(<-ch)

	fmt.Println("in main: blocking 2sec occurs immed after first thread, since 5sec already passed")
	fmt.Println(<-ch2)

	// fmt.Println("in main: write to channel should deadlock since same thread")
	// ch <- 3

	fmt.Println("in main: end")
}
