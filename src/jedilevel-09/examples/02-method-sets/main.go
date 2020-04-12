package main

import "fmt"

func main() {
	p1 := person{"Abc"}
	saySomething(&p1)
}

type person struct {
	Name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello from", p.Name)
}

func saySomething(h human) {
	h.speak()
}
