package main

import (
	"fmt"
	"time"
)

func readThread(c chan int) {
	fmt.Println("in thread: block read until write occurs")
	msg := <-c
	fmt.Println("in thread:", msg, "read")
}

func main() {
	ch := make(chan int)
	fmt.Println("in main: call read thread")
	go readThread(ch)
	fmt.Println("in main: delay write by 5 seconds")
	time.Sleep(time.Second * 5) // delay to demonstrate write delay
	fmt.Println("in main: write & unblock")
	ch <- 5
	time.Sleep(time.Second * 1) // give thread some time to exit
	fmt.Println("in main: end")
}
