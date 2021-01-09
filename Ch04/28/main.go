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
	go func(ch chan int, wg *sync.WaitGroup) {
		// use if statement to determine whether channel had been closed or not for printing received message
		if meg, ok := <- ch; ok {
			fmt.Println(meg)
		}
		wg.Done()
	} (channel, waitGroup)

	// create second Goroutine to will include send integer message on channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// ch <- 0 // test if send message
		close(ch) // test close channel before send message
		wg.Done()
	} (channel, waitGroup)

	// finally set waitGroup to wait
	waitGroup.Wait()
}