package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("This is my first Func")
		wg.Done()
	}()

	go func() {
		fmt.Println("This is my second Func")
		wg.Done()
	}()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("About to Exit !")
}
