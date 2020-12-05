package main

import (
	"fmt"
)

// Test for try Channels without Go routines (Passing message to channel then Print it)
func main() {
	// create channel for integer messages
	channel := make(chan int)

	channel <- 42
	fmt.Println(<- channel)
}