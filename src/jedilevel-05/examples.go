package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavour []string
}

func main() {
	loop_struct()
	keep_in_map()
	embeding_types()
	annonimous_struct()
}

// ==== Example 1 ====
// Create struct and apply a for loop

func loop_struct() {
	fmt.Println("=== Simple Struct Storage ===")
	p1 := person{
		first:      "data",
		last:       "psycho",
		favFlavour: []string{"vanilla", "orange"},
	}

	p2 := person{
		first:      "String",
		last:       "Rider",
		favFlavour: []string{"smokey", "cherry"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for k, v := range p1.favFlavour {
		fmt.Println(k, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for k, v := range p2.favFlavour {
		fmt.Println(k, v)
	}
}

// Store custom types in a map
func keep_in_map() {
	fmt.Println("=== Map Storage ===")
	p1 := person{
		first:      "data",
		last:       "psycho",
		favFlavour: []string{"vanilla", "orange"},
	}

	p2 := person{
		first:      "String",
		last:       "Rider",
		favFlavour: []string{"smokey", "cherry"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

}

// === Embeding A Struct to Another Struct ===
func embeding_types() {
	fmt.Println("=== Embeding Type ===")
	type vehicle struct {
		door  int
		color string
	}

	type truct struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	// creating value
	t := truct{
		vehicle: vehicle{
			door:  2,
			color: "yellow",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			door:  4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}

// Creating Annonimous Struct
func annonimous_struct() {
	fmt.Println("=== Annonimous Struct ===")
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"MoneyPenny": 555,
			"Q":          777,
			"Psycho":     675,
		},
		favDrinks: []string{
			"Coke",
			"Water",
		},
	}
	fmt.Println(s)
}
