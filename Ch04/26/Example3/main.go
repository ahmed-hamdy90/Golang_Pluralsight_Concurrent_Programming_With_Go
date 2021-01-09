package main

import (
	"fmt"
	"sync"
)

func main() {
	// create wait group and add how much to wait
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	// create channel with integer messages
	channel := make(chan int)

	// create first Goroutine to will include Print received message from channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<- ch)
		// close channel after receive message from channel
		// close(ch) // this line will throw Panic as only you can't close receiving channel type
		wg.Done()
	} (channel, waitGroup)

	// create second Goroutine to will include send integer message on channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	} (channel, waitGroup)

	// finally set waitGroup to wait
	waitGroup.Wait()
}