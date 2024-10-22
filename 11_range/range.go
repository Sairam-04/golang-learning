package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

	map1 := map[string]string{"name": "sairam", "age":"22"}
	for k, v := range map1{
		fmt.Println(k,v)
	}

	// unicode code point rune
	// first element i is the starting byte of rune
	// characters more than 255 takes more than one byte
	// emoji takes 4 bytes
	 for i,c := range "🐼golang"{
		fmt.Println(i,c)
	 }
}