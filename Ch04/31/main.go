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
	// use waitGroup instead of time sleep code (preferred use pointer for passing it)
	waitGroup := &sync.WaitGroup{}
	// use RWMutex to protect write/read data during Goroutine (preferred use pointer for passing it)
	// use read/write mutex on large code as there multi-reader and only one writer
	mutex := &sync.RWMutex{}
	// create two channels for query from cache and database, which will sent query Book result
	cacheChannel := make(chan Book)
	databaseChannel := make(chan Book)

	for counter:= 0; counter < 10; counter++ {
		id := random.Intn(10) + 1
		// Must define how much wait[daily] need(combine or separated => waitGroup.Add(1) twice)
		waitGroup.Add(2)
		// use Goroutine to querying from cache (passing waitGroup)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if book, ok := queryCache(id, m); ok {
				ch <- book
			}
			// Must define complete point for waitGroup
			wg.Done()
		}(id, waitGroup, mutex, cacheChannel)
		// use Goroutine to querying from database (passing waitGroup)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if book, ok := queryDatabase(id, m); ok {
				ch <- book
			}
			// Must define complete point for waitGroup
			wg.Done()
		}(id, waitGroup, mutex, databaseChannel)
		// use Goroutine as third one to receive message from both Cache and Database channels
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case book := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(book)
				// must receive too message from Database channel as it always send message
				// if not receive it message will throw deadlock error
				<-dbCh
			case book := <-dbCh:
				fmt.Println("from database")
				fmt.Println(book)
			}
		}(cacheChannel, databaseChannel)
		// add sleep time for wait Goroutine complete
		time.Sleep(150 * time.Millisecond)
	}
	// Must tell waitGroup to wait
	waitGroup.Wait()
}

// querying Book instance by id which had been cached before
func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	// set mutex lock and protect any data (any code will block under mutex.RLock line until mutex.RUnLock)
	// use read mutex lock/unlock methods as this line define reader
	m.RLock()
	// check if book already cached before or not
	cachedBook, ok := cache[id]
	m.RUnlock()

	return cachedBook, ok
}

// querying Book instance by id from Database
func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	// make simple sleep time
	time.Sleep(100 * time.Millisecond)
	// loop on saved Books to search for book which we need
	for _, book := range books {
		if book.ID == id {
			// set mutex lock and protect any data (any code will block under mutex.Lock line until mutex.UnLock)
			// use ordinal lock/unlock methods as this line define writer(the only one)
			m.Lock()
			// cache find Book instance
			cache[id] = book
			m.Unlock()

			return book, true
		}
	}

	return Book{}, false
}