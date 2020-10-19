package main

/*
The capacity, in number of elements, sets the size of the buffer in the channel.
If the capacity is zero or absent, the channel is unbuffered
and communication succeeds only when both a sender and receiver are ready.

Otherwise, the channel is buffered and communication succeeds without blocking
if the buffer is not full (sends) or not empty (receives).

A nil channel is never ready for communication.
*/

import (
	"fmt"
)

func readFromChannel(ch <-chan int, value *int) {
	inval := <-ch
	value = &inval
}

func main() {
	channelread := make(<-chan int) // create read channel only

	var aptr *int
	b := 555

	aptr = &b
	go readFromChannel(channelread, aptr)
	fmt.Println(*aptr)

	b = 666

	go readFromChannel(channelread, aptr)
	fmt.Println(*aptr)
}
