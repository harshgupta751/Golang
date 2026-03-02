package main

import "fmt"

func main() {

	myNumber := 53
	var ptr *int = &myNumber

	fmt.Println("ptr= ", ptr);
	fmt.Println("Value at Ptr", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Updated value", myNumber)


}