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

	// create first Goroutine to will include Print received messages from channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// First solution for receive multi messages form channel using for loop iterator, use for loop iterator too
		// But this way we must know messages number will be sent
		//for counter := 0; counter < 10; counter++ {
		//	fmt.Println(<- ch)
		//}

		// Second solution is use range keyword to loop on sent messages [Best solution]
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	} (channel, waitGroup)

	// create second Goroutine to will include send integer messages on channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// in case we need to send multi messages using for loop iterator
		for counter := 0; counter < 10; counter++ {
			ch <- counter
		}
		// you must close sender channel after for loop iterator to prevent Deadlock error
		// especial in case use range keyword to receive sent messages
		close(ch)
		wg.Done()
	} (channel, waitGroup)

	// finally set waitGroup to wait
	waitGroup.Wait()
}