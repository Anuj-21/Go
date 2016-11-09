/*do not provide any inbuilt method to increase size of it dynamically or get a sub-array of its own.Slices covers this limitation*/
package main

import "fmt"

func main() {
	//var slice []int // a slice of unspecified size
	// slice == []int{0,0,0,0,0}
	var slice = make([]int, 3, 5) // a slice of length 3 and capacity 5
	printSlice(slice)

	// print the original slice
	fmt.Println("slice ==", slice)

	// print the sub slice starting from index 1(included) to index 4(excluded)
	fmt.Println("slice[1:4] ==", slice[1:4])

	// missing lower bound implies 0
	fmt.Println("slice[:3] ==", slice[:3])

	// missing upper bound implies len(s)
	fmt.Println("slice[4:] ==", slice[4:])

	slice1 := make([]int, 0, 5)
	printSlice(slice1)

	// print the sub slice starting from index 0(included) to index 2(excluded)
	number2 := slice[:2]
	printSlice(number2)

	// print the sub slice starting from index 2(included) to index 5(excluded)
	number3 := slice[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
