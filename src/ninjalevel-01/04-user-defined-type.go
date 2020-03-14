package main

import "fmt"

type mytype int

var myvar mytype

func main() {
	fmt.Println(myvar)
	// %T for printing the type
	fmt.Printf("%T\n", myvar)
	myvar = 5
	fmt.Println(myvar)
}
