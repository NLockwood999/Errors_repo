package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func areacircle(radius float64) (float64, error) {
	//check if radius is less than 0
	if radius < 0 {
		err := errors.New("You cannot have a negative radius")
		return -1, err
	}
	result := (math.Pi * math.Pow(radius, 2))
	return result, nil
}

func main() {
	r := 20.0
	result, err := areacircle(float64(r))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(math.Round(result))
}
