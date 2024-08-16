package main

import (
	"fmt"
)

// Data types:
// string
// bool
// int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
// float32, float64
// complex64, complex128

func pracInt() {
	var x, y int = 1, 3
	a, b := 10, 20
	fmt.Println(x, y, a, b)

}

func pracString() {
	var firstname string = "Brian"
	lastname := "Kam"
	var name = "Anna Mori"
	fmt.Println(name)
	fmt.Println(firstname, lastname)
	fmt.Printf("%s\n", firstname+" "+lastname)
}

func main() {
	// pracString()
	pracInt()
}

// go run "file path"s
