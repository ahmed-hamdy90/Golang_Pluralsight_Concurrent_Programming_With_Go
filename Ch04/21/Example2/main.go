package main

import (
	"fmt"
)

// Test try Channels without Go routines (Print message before passing it to channel)
func main() {
	// create channel for integer messages
	channel := make(chan int)

	fmt.Println(<- channel)
	channel <- 42
}
