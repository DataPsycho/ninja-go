package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print the memory location of a value
	fmt.Println("==== Exercise 1 ====")
	a_int := 32
	fmt.Printf("I am a %T My address is %v.\n", a_int, &a_int)
	// Exercise 2
	p1 := person{"Data Nurd"}
	fmt.Printf("My old name is : %v.\n", p1.name)
	change_name(&p1, "Data Psycho")
	fmt.Printf("My new name is : %v.\n", p1.name)

}

// ==== Exercise 2 =====
// Create a struct with person
// Cerate a mothod for that struct is
// capable ot changing the name of the person using pointer
type person struct {
	name string
}

func change_name(p *person, n string) {
	(*p).name = n
}
