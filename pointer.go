package main

import (
	"fmt"
)

func main() {
	var a int = 20 //actual variable declaration
	var ip *int    //pointer variable declaration
	var ptr *int

	ip = &a //storing address in the pointer

	fmt.Printf("Address of a variable : %x\n", &a)

	//address store in pointer variable
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	//access the value using the pointer

	fmt.Printf("Value of *ip variable: %d\n", *ip)
	fmt.Printf("Value of *ip variable: %d\n", ptr)
}
