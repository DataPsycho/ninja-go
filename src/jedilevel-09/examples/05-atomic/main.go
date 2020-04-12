package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup

	// Setting up Atomic Incrementer
	var incrementer int64

	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value:", incrementer)
}
