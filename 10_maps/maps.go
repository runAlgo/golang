package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map

	//          key:type value:type
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["phone"])

	// IMP: if key does not exists in the map then it returns zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 23
	// m["roll"] = 10

	// delete specific element from map
	// delete(m, "age")
	// Delete all element from map
	// clear(m)

	// fmt.Println("Length of map:", len(m))
	// fmt.Println(m)



	// Another method to create Map: it is used when we know Exact data
	m1 := map[string]int{"user": 1, "user2": 2, "user3": 3}
	m2 := map[string]int{"user": 1, "user2": 2, "user3": 5}

	fmt.Println(maps.Equal(m1, m2))
	value, ok := m1["user"]
	fmt.Println(value)
	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("not ok")
	}

}