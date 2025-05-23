package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices = Dynamic arrays
	// 	- Slices are a flexible and powerful way to work with collections of data in Go.
	// 	- They are built on top of arrays and provide a more convenient and efficient way to handle collections of elements.

	// uninitializes scile is nill
	// var nums []int
	// fmt.Println(nums==nil) // nil
	// fmt.Println(len(nums)) // 0

	// var nums = make([]int,2,5) // 2 is length(it will store 2 zeros entry at the beginning of nums) and 5 is capacity
	// fmt.Println(nums) // [0 0]
	// fmt.Println(nums==nil) // false
	// fmt.Println(len(nums)) // 2
	// // capacity -> maximum no. of elements that can be stored in the slice
	// fmt.Println(cap(nums)) // 5
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// // now if we append anything it will increase the capacity automatically
	// // it will double the capacity of prevoius capacity
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// fmt.Println(nums) // [0 0 1 2 3 4 5]
	// fmt.Println(len(nums)) // 7
	// fmt.Println(cap(nums)) // 10

//  Another way of defining slice
	// nums := []int{}
var nums = make([]int, 0, 5) // 0 is length(it will store 0 zeros entry at the beginning of nums) and 5 is capacity
	nums = append(nums, 1)
var nums2 = make([]int, len(nums), 5) 

	copy(nums2, nums) // copy the elements of nums to nums2
	fmt.Println(nums) // [1 0 0 0 0]
	fmt.Println(nums2) // [1 0 0 0 0]

	// slice operator
	nums4 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums4[0:2]) // [1 2]
	fmt.Println(nums4[:2])  // [1 2]


	//  slice 

	nums3 := []int{1,3,4,66,89,9}

	fmt.Println(slices.Equal(nums, nums3)) // true

var nums5 = [][]int{{1, 2,3}, {4, 5,6}}

	fmt.Println(nums5) // [[1 2] [3 4] [5 6]]

}