package main

import (
	"fmt"
	"math"
)

func main() {
	n := foo()
	x, s := bar()
	fmt.Println("=== Execusiion of exercise 1 ===")
	fmt.Println("A number", n)
	fmt.Println("My flight to Protugal:", x, s)

	// Exercise 2
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total_variadic := sum_variadic(ii...)
	fmt.Println("=== Execusiion of exercise 2 ===")
	fmt.Println("The total from varaidic is:", total_variadic)
	total_slice := sum_slice(ii)
	fmt.Println("The total from slice is:", total_slice)

	// Exercise 3
	fmt.Println("=== Execusiion of exercise 3 ===")
	defer defer_func()
	before_func()

	// Exercise 4
	fmt.Println("=== Execusiion of exercise 4 ===")
	me := person{"Data", "Psycho", 29}
	me.speak()

	// Exercise 5
	fmt.Println("=== Execusiion of exercise 5 ===")
	a_circle := circle{5.2}
	a_square := square{7}
	info(a_circle)
	info(a_square)

	// Exercise 6
	// annonymous func
	fmt.Println("=== Execusiion of exercise 6 ===")
	func() {
		x := 0
		for i := 0; i < 100; i++ {
			x += i
		}
		fmt.Printf("Adding 0 to 99 in a anonymous function %v.\n", x)
	}()
	// Exercise 7
	fmt.Println("=== Execusiion of exercise 7 ===")
	f := assigned_func
	fmt.Printf("I am a assigned function with type %T.\n", f)
	f()

	// Exercise 8
	fmt.Println("=== Execusiion of exercise 8 ===")
	number_1 := 5
	rf := return_func()
	rf(number_1)

	// Exercise 9
	fmt.Println("=== Execusiion of exercise 9 ===")
	cf := callback_func(called_func, []int{1, 2, 3, 4})
	fmt.Printf("Result from the call back func %v.\n", cf)

}

// ===== Exercise 1 =====
// create a func return a int
func foo() int {
	return 42
}

// create a func return a string and int
func bar() (int, string) {
	return 1971, "Fly Tap"
}

// ==== Exercise 2 ====
// create a func which take a variadic parameter
// takes a slice and return a sum of the slice
func sum_variadic(xi ...int) int {
	total := 0
	for _, k := range xi {
		total += k
	}
	return total
}

// take a slice and return sum
func sum_slice(slice []int) int {
	total := 0
	for _, k := range slice {
		total += k
	}
	return total
}

// ==== Exercise 3 ====
// Use defer to run a func after the execution or
// any other func
func defer_func() {
	fmt.Println("Running Defer!")
}

func before_func() {
	fmt.Println("Running Before Defer !")
}

// ==== Exercise 4 ====
// An uds with identifier person
// with the fields first, last, age
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Introducing Myself !!!")
	fmt.Printf("I am %v %v, I am %v year old. \n", p.first, p.last, p.age)
}

// ==== Exercise 5 ====
// A square, A circle a method area for square and circle
// A type shape with the type interface as anything that has area method
// create func info which take type shape and print area
type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("The area of %T is %v.\n", s, s.area())
}

// ==== Exercise 6 ====
// Create a annonymous function
// annonumous function applied on main function

// ==== Exercise 7 ====
// assign a func to a variable and call it

func assigned_func() {
	x := 0
	for i := 0; i < 100; i++ {
		x += i
	}
	fmt.Printf("Adding 0 to 99 in a anonymous function %v.\n", x)
}

// ==== Exercise 8 ====
// Create a func which return a func, assigned a func to a variable
// the call the func

func return_func() func(x int) {
	return func(x int) {
		if x%2 == 0 {
			fmt.Printf("%v is a even number.\n", x)
		} else {
			fmt.Printf("%v is a odd number.\n", x)
		}
	}
}

// ==== Exercise 9 ====
// callback called_func is used int callback_func as a parameter
func called_func(xi []int) int {
	n := 0
	for _, k := range xi {
		n += k
	}
	return n
}

func callback_func(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
