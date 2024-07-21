package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	lock    = sync.Mutex{}
)

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	lock.Lock()
	counter++
	fmt.Printf("Counter: %d\n", counter)
	lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go Worker(&wg)
	}

	wg.Wait()
	fmt.Println("Main finish.")
}
