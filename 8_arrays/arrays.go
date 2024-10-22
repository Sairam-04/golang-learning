package main

import "fmt"

// Arrays are numbered sequence of specific length
// fixed size Memory optimization
//  Costant time access 
func main(){
	var nums [4]int

	// adding array elements
	nums[0] = 100
	fmt.Println("nums[0]:", nums[0])
	fmt.Println("nums :",nums)
	// array Length
	fmt.Println("Length of Array", len(nums))

	// to declare in single line
	numsArr := [3]int{1,2,300}
	fmt.Println(numsArr)

	// 2d Array
	nums2D := [2][2]int{{1,2},{2,3}}
	fmt.Println(nums2D)
}