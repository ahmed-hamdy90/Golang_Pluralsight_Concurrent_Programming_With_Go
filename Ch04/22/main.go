package main

import (
	"fmt"
	"sync"
)

func main() {
	// create wait group and add how much to wait
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	// create Buffered channel with integer messages which determine number of messages will be receive,
	// any extra messages add to channel will discard them
	channel := make(chan int, 1)

	// create first Goroutine to will include Print received message from channel
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<- ch)
		wg.Done()
	} (channel, waitGroup)

	// create second Goroutine to will include send integer message on channel
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		wg.Done()
	} (channel, waitGroup)

	// finally set waitGroup to wait
	waitGroup.Wait()
}