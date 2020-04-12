package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup

	// Setting up waitGroup
	incrementer := 0
	const gs = 100
	wg.Add(gs)

	// Define Mutex
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// Apply lock
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			// Apply unlock
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("GOroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Coun:", incrementer)
}
