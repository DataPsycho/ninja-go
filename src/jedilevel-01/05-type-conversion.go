package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Println(x)
	// %T for printing the type
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T", y)
	fmt.Println("\n")
}
