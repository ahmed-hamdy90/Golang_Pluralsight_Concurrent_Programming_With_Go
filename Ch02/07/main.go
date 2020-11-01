package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book {}
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// main function for program
func main() {
	for counter:= 0; counter < 10; counter++ {
		id := random.Intn(10) + 1
		// use Goroutine to querying from cache
		go func(id int) {
			if book, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(book)
			}
		}(id)
		// use Goroutine to querying from database
		go func(id int) {
			if book, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(book)
			}
		}(id)
		// make simple sleep time (to enable Go routine execute)
		time.Sleep(150 * time.Millisecond)
	}
	// another simple sleep time (to catch the last go routine before program exit)
	time.Sleep(2 * time.Second)
}

// querying Book instance by id which had been cached before
func queryCache(id int) (Book, bool) {
	// check if book already cached before or not
	cachedBook, ok := cache[id]

	return cachedBook, ok
}

// querying Book instance by id from Database
func queryDatabase(id int) (Book, bool) {
	// make simple sleep time
	time.Sleep(100 * time.Millisecond)
	// loop on saved Books to search for book which we need
	for _, book := range books {
		if book.ID == id {
			// cache find Book instance
			cache[id] = book

			return book, true
		}
	}

	return Book{}, false
}