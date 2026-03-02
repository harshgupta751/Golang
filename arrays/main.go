package main

import "fmt"

func main() {

	var myList [4]string
	myList[0] = "Harsh"
	myList[1] = "Ramesh"
	myList[3] = "Suresh"

	fmt.Println(myList)
	fmt.Println(len(myList))

	var vegList = [3]string{"Cabbage", "Potato"}

	fmt.Println(vegList)
fmt.Println(vegList[0])
fmt.Println(vegList[1])
fmt.Println(vegList[2])
	fmt.Println(len(vegList))

}