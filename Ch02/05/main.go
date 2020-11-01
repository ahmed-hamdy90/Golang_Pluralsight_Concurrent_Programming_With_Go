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
		// first querying from cache
		if book, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(book)
			continue
		}
		// then querying from database
		if book, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(book)
			continue
		}
		// otherwise
		fmt.Printf("Book not found with id: '%v'", id)
		// make simple sleep time
		time.Sleep(150 * time.Millisecond)
	}
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