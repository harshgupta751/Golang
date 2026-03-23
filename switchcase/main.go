package main

import (
	"fmt"
 "math/rand"
	"time"
)


func main() {

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
fmt.Println("Dice value is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, You can open")
	case 2: 
		fmt.Println("You can move 2 spot")
	case 3, 4: 
		fmt.Println("You can move 3 or 4 spot")
		fallthrough
	// case 4: 
	// 	fmt.Println("You can move 4 spot")
	case 5: 
		fmt.Println("You can move 5 spot")
	case 6: 
		fmt.Println("You can move to spot 6 and roll dice again")
	default:
		fmt.Println("Invalid diceNumber")
	}

}