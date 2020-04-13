package main

import (
	"encoding/json"
	"fmt"
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

	bs, err := toJson(p1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

// toJson need to return an error
func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: %v", err)
	}
	return bs, nil
}
