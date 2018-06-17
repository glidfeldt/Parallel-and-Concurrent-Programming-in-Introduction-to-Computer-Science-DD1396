package main

import (
	"fmt"
	"sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {

	waitG := new(sync.WaitGroup)
	waitG.Add(11)
	ch := make(chan int)
	go Print(ch, waitG)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	waitG.Wait()
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, waitG*sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
		waitG.Done()
	}
}
