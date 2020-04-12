package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Exercise 1
	fmt.Println("==== Example 1 ====")
	u1 := user{"Data Psycho", 29}
	u2 := user{"Data Nerd", 32}
	u3 := user{"String Rider", 30}
	users := []user{u1, u2, u3}
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The Marshalled Data:", string(bs))

	// Exercise 2
	fmt.Println("==== Example 3 ====")
	s := string(bs)
	var unm_values []unmarsUser
	err2 := json.Unmarshal([]byte(s), &unm_values)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("Go Native Data:", unm_values)
	// Exercise 3
	eu1 := enccodeUser{"James", "Bond", 32, []string{"Shaken, not stirred"}}
	eu2 := enccodeUser{"Miss", "Moneypenny", 27, []string{"James, it is soo good to see you"}}
	eu3 := enccodeUser{"M", "Hmmmm", 54, []string{"Oh, James. You didn't."}}
	eu_values := []enccodeUser{eu1, eu2, eu3}
	err3 := json.NewEncoder(os.Stdout).Encode(eu_values)
	if err3 != nil {
		fmt.Println(err)
	}

	fmt.Println("==== Example 4 ====")
	// === Exercise 4 ===
	xi := []int{4, 1, 12, 11}
	// sort package need to be imported
	fmt.Println("Before sort: ", xi)
	sort.Ints(xi)
	fmt.Println("After sort: ", xi)

	fmt.Println("==== Example 5 ====")
	// === Exercise 5 ===
	// Sort the eu_encode from Exercise 3 by age
	unsorted_values := []enccodeUser{eu1, eu2, eu3}
	sort.Sort(ByAge(unsorted_values))
	fmt.Println(unsorted_values)

}

// === Exercise 1 ===
// A slice of type user to json using Marshal
type user struct {
	Name string
	Age  int
}

// === Exercise 2 ===
// Unmarshal
// json string to Composite data type
type unmarsUser struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

// === Exercise 3 ===
//  encode to JSON the []user sending the results to Stdout.
type enccodeUser struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// === Exercise 4 ===
// Sorting a slice of string and int
// Direct implemented in main

// === Exercise 5 ===
// Directly Implemented in Main
type ByAge []enccodeUser

// implementing the 3 methods to ByAge type
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
