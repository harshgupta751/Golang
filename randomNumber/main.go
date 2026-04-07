package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main(){

var num1 int = 61
var num2 float64 = 7.1

fmt.Println("The Sum is :", num1 + int(num2))


	//random from math/rand
	// rand.Seed(time.Now().UnixNano())
	// randomNum := rand.Intn(6) + 1 
	// fmt.Println(randomNum)

	//random from crypto/rand
	n, _ := rand.Int(rand.Reader, big.NewInt(6))
	fmt.Println(n.Int64() + 1)



}