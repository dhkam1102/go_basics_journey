package main

import (
	"math/rand"
)

// value is passed to funciton, its copy is passed
func makeMeOlder(age int) {
	age += 10
}

// need to specify the return type
func getRandomCar() string {
	brands := []string{
		"KIA", "BMW", "TESLA", "FORD",
	}
	var random_number int = rand.Intn(4)
	return brands[random_number]
}

// when multiple retunr type do:
// func getRandomCar() (string, int)
