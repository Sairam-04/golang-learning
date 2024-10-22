package main

import "fmt"

// variadic functions are those where we can pass any number of parameters 

func sum(nums ...int) int{
	total := 0
	for _, num := range nums{
		total+=num
	}
	return total 
}
func main() {
	nums := []int{10,11,12,13,14,15}
	fmt.Println("nums :", nums, "\nSum :", sum(nums...))
	fmt.Println("sum(1,2,3,4,5) =", sum(1,2,3,4,5))

}