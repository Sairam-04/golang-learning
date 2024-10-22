package main

import "fmt"

func changeNum(num *int) {
	*num = 100
}

func main() {
	num := 10
	fmt.Println("Before Changing", num)
	changeNum(&num)
	fmt.Println("After Changing", num)
}