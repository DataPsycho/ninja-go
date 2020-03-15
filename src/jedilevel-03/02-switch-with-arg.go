package main

import "fmt"

func main() {
	favSport := "badminton"
	// switch will evaluae if case argument is true
	// in that case it is badminton
	switch favSport {
	case "skiing":
		fmt.Println("Go to mountain")
	case "badminton":
		fmt.Println("Go to branik")

	}
}
