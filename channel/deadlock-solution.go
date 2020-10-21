package main

/*
If a Goroutine is sending data on a channel, then it is expected
that some other Goroutine should be receiving the data.
If this does not happen, then the program will panic at runtime with Deadlock.
*/

import "fmt"

func routine(cha chan<- int, i int) {
	cha <- i // sending data on channel
}

func main() {
	ch := make(chan int)
	go routine(ch, 5)
	fmt.Println(<-ch) // main goroutine exists to receive it
}
