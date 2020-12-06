package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create wait group and add how much to wait
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	// create channel with integer messages
	channel := make(chan int)

	// create first Goroutine to will include Print received message from channel
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<- ch)
		ch <- 27 // as this passing channel is bi-directional so we can send and receive messages
		wg.Done()
	} (channel, waitGroup)

	// create second Goroutine to will include send integer message on channel
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		// Must setting sleep time to wait until sending message will be receive
		time.Sleep(5 * time.Millisecond)
		fmt.Println(<- ch)
		wg.Done()
	} (channel, waitGroup)

	// finally set waitGroup to wait
	waitGroup.Wait()
}