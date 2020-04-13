package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

// Implement Error Inteface
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v; %v.", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.10)
	if err != nil {
		log.Println(err)
	}
}

// Function to generate Error
func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("I need more Tea!")
		return 0, sqrtError{"50.22 N", "99.23 W", e}
	}
	return 42, nil
}
