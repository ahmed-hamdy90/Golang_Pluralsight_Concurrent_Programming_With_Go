package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book {}
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// main function for program
func main() {
	// use waitGroup instead of time sleep code
	waitGroup := &sync.WaitGroup{}
	for counter:= 0; counter < 10; counter++ {
		id := random.Intn(10) + 1
		// Must define how much wait[daily] need(combine or separated => waitGroup.Add(1) twice)
		waitGroup.Add(2)
		// use Goroutine to querying from cache (passing waitGroup)
		go func(id int, wg *sync.WaitGroup) {
			if book, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(book)
			}
			// Must define complete point for waitGroup
			wg.Done()
		}(id, waitGroup)
		// use Goroutine to querying from database (passing waitGroup)
		go func(id int, wg *sync.WaitGroup) {
			if book, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(book)
			}
			// Must define complete point for waitGroup
			wg.Done()
		}(id, waitGroup)
		// make simple sleep time (to enable Go routine execute)
		time.Sleep(150 * time.Millisecond)
	}
	// Must tell waitGroup to wait
	waitGroup.Wait()
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