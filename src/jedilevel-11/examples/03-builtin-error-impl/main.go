package main

import "fmt"

type customErr struct {
	info string
}

// Implementing error interface
func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v.\n", ce.info)
}

func main() {
	c1 := customErr{"Need more tea !"}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}
