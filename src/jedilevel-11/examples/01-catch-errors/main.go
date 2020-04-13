package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name        string
	Age         int
	Description string
}

// Exerice:
// Catch the error if any all the time
func main() {
	p1 := person{"Data Psycho", 29, "I do not have life!"}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
