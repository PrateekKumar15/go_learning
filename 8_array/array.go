package main

import "fmt"

func main() {
	//  numbered sequence of specific length
	//  in case of num array it stores value to be zero until we assign any value to it
	//  in case of bool array it stores value to be false until we assign any value to it
	//  in case of string array it stores value to be empty string until we assign any value to it
	var nums [4]int

	fmt.Println(len(nums))
	nums[0] = 1
	fmt.Println(nums[0])
	nums[1] = 2
	fmt.Println(nums) 

num := [3]int{1, 2, 3}
	fmt.Println(num)

	twoDArr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDArr) 
//  fixed size , that is predictable
//  Memory Optimization
// constant time access

}