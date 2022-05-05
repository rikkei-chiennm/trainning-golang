package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		// multiple go routine can read(not write) at a time by acquiring the lock.
		go sayHello()
		m.Lock()
		//  Locks the mutex, m. If the lock is already in use, call goroutine blocks until the mutex is available.
		go increment()
	}
	wg.Wait()
}
func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	// Unlocks the mutex, m. A run-time error is thrown if m is not already locked.
	wg.Done()
}
