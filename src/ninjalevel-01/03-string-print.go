package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// For sprint you need location
	// identifier
	// Sprint + Format: Sprintf
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
