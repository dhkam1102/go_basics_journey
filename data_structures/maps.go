package data_structures

import (
	"fmt"
)

type hello struct {
	first, second int
}

func pracMap() {
	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1
	// fmt.Println(m)

	var hi = map[string]hello{
		"first":  hello{1, 2},
		"second": {2, 3},
	}
	hi["thrid"] = hello{4, 5}
	// fmt.Println(hi)

	random := make(map[string]int)
	random["first"] = 5
	random["second"] = 10
	delete(random, "first")
	fmt.Println(random)
	value, ok := random["second"]
	fmt.Println("The value:", value, "Present?", ok)
	v := random["second"]
	fmt.Printf("value: %d \n", v)
}

// func main() {
// 	fmt.Println("map")
// 	pracMap()
// }
