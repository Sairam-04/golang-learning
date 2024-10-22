package main

import "fmt"

// Slices are dynamic in size
//  It is the most  used construct in go
// It has all useful methods for slice manipulation
func main(){
	// unintialized slice is nil
	var nums []int
	fmt.Println("Length :",len(nums), "nums == nil", nums == nil)

	// make function to intialize the slice
	// It is the initial size (2)
	//  second argument is the intitial cap
	// third parameter is capacity
	var sl = make([]int, 2, 10)
	fmt.Println("len(sl)", len(sl), "cap(sl)", cap(sl))

	// slices resizes itself when additional elements added in the slice
	sl = append(sl, 100)
	sl = append(sl, 200)
	fmt.Println("sl",sl)

	// another way of making slice
	// slice1 := []int{}

	// Methods
	/*
		1. copy()
		2. append()
		3. slices.Equal(slice1, slice2)

	*/
	var nums1 = make([]int, 0, 5)
	nums1 = append(nums1, 2)
	var nums2 = make([]int, len(nums1))
	copy(nums2, nums1)
	fmt.Println("copy",nums1, nums2)
}