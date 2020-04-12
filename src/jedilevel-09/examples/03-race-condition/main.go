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

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
		fmt.Println("GOroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Coun:", incrementer)
}
