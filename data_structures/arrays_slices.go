package main

import (
	"fmt"
)

func arrayFormat() {
	// var hello [2]string
	// hello[0] = "hello"
	// hello[1] = "world"
	// fmt.Println(hello[0], hello[1])
	// fmt.Println(hello) // prints [hello world]

	// nums := [6]int{1, 2, 3, 4, 5, 6}
	// fmt.Println(nums)

	// var s []int = nums[1:4]
	// fmt.Println(s)

	// NOTE: Slices are references to array
	// names := [4]string{
	// 	"Brian",
	// 	"Anna",
	// 	"Tony",
	// 	"Martha",
	// }

	// a := names[0:2]
	// b := names[1:3]

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// NOTE Slices are dynamically sized
	// slice1 := []int{1, 2, 3, 4}
	// slice1 = append(slice1, 5, 6, 7)
	// fmt.Println(slice1)

	// NOTE Slices are nil with no value
	// var s []int
	// if s == nil {
	// 	fmt.Println("nil")
	// }

	// NOTE: make function allocates 0
	// 		parameter: type of slice, len, cap
	a := make([]int, 3, 5)
	fmt.Println(a)

}

func main() {
	fmt.Println("array")
	arrayFormat()

}
