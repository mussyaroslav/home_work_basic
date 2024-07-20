package main

import (
	"sync"
	"testing"
)

func ResetCounter() {
	lock.Lock()
	counter = 0
	lock.Unlock()
}

func TestWorker(t *testing.T) {
	ResetCounter()

	var wg sync.WaitGroup
	numWorkers := 500

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(&wg)
	}

	wg.Wait()

	if counter != numWorkers {
		t.Errorf("Expected counter to be %d, but got %d", numWorkers, counter)
	}
}
