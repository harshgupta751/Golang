package main

import "fmt"

func main() {

	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Println(fruitList)
	fmt.Printf("Type of FruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Guava", "Mango")
	fmt.Println(fruitList)
	fmt.Printf("Type of FruitList is %T\n", fruitList)

	// fruitList= append(fruitList[1:])
	fruitList= append(fruitList[0:3])
	
	fmt.Println(fruitList)

	var myList []string
	fmt.Printf("Type of MyList is %T\n", myList)
	fmt.Println(myList)
	myList= append(myList, "Hi", "Golang")
	fmt.Printf("Type of MyList is %T\n", myList)
	fmt.Println(myList)

	 highScores := make([]int, 3 )
	highScores[0]= 2;
	highScores[1]=53
	highScores[2]=93
	fmt.Println(highScores)

	highScores = append(highScores, 1, 100)
	fmt.Println(highScores)


}