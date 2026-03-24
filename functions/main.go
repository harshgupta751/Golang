package main

import "fmt"

func main() {

	val1, val2 := adder(4,6);
	fmt.Println(val1);
	fmt.Println(val2)
fmt.Println(proAdder(12,2,6,3,5))

}

func proAdder(values ...int) int {
total:= 0
for _, val:= range values {
	total += val
}
return total

}


func adder(val1 int, val2 int) (int, string) {

	return val1 + val2, "Hello"
}