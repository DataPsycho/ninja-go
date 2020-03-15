package main

import "fmt"

func main() {
	arrayExample()
	sliceExample()
	mapExample()
}

func arrayExample() {
	fmt.Println("========= Array Example ======")
	x := [5]int{1, 8, 6, 4}
	// ======== Operations ===========
	// Loop to print each value
	for i, v := range x {
		fmt.Println(i, v)
	}
	// Printing types, the last element will
	// return 0 as it is not assigned
	fmt.Printf("%T\n", x)

}

func sliceExample() {
	fmt.Println("========= Slice Example ======")
	x := []int{1, 2, 8, 3, 9, 4, 5, 7}
	fmt.Printf("%T\n", x)
	// ======= Operation ========
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("== Appending a new slice ==")
	y := []int{45, 78, 8}
	x = append(x, y...)
	fmt.Println(x)

	fmt.Println("== Slice with Make ==")
	district := make([]string, 6, 6)
	district = []string{"Dhaka", "Rajshahi", "Rangpur"}
	fmt.Println(district)

	for i := 0; i < len(district); i++ {
		fmt.Println(i, district[i])
	}
}

func mapExample() {
	fmt.Println("======== Map Example =========")
	m := map[string][]string{
		"psycho_data":  []string{"metal"},
		"rock_monger":  []string{"hard rock"},
		"death_monger": []string{"death metal"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the rocrd for\n", k)
		for i, v2 := range v {
			fmt.Printf("\t%v, %v", i, v2)
		}
	}
	fmt.Println("\n ==")

	fmt.Println("====== Adding A Record ======")
	m["neo_monger"] = []string{"new classic"}
	fmt.Println(m)
	fmt.Println("=== Deleting neo_monger ====")
	if _, ok := m["neo_monger"]; ok {
		delete(m, "neo_monger")
		fmt.Println(m)
	}
}
