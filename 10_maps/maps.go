package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	fmt.Println(m[1])

	// if key doesnot not exists in map then it returns zero value
	// length of map
	fmt.Println("len(m)", len(m), "map m:", m,"m[100]", m[100])
	/*
		Methods
		1. delete(map, key)
		2. clear(map)
		3. maps.Equal(map1, map2)
	*/

	// Initializing map without make
	m1 := map[string]int{"price":40,"phones":2}
	m2 := map[string]int{"price":40,"phones":2, "quantity":10}
	fmt.Println(m1)

	v, ok := m1["name"]
	//  the second parameter that is returned is bool if key exists it is true else false
	if ok{
		fmt.Println("Key is present", v)
	}else{
		fmt.Println("Key is not present", v)
	}
	fmt.Println(maps.Equal(m2, m1))
}