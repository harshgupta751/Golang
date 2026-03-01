package main

import "fmt"                  // formatting library package

var amount= 53

const AccessTokenExpiry = 60      // varible starts with capital means Public variable

func main (){

fmt.Print("Hello ")
fmt.Print("World")
fmt.Println()

	//
var username string = "Harsh Gupta"

fmt.Println("My username is :", username)

fmt.Printf("Type of username is %T\n", username)

//
var age int = 93

fmt.Println("My age is :", age)

fmt.Printf("Type of age is %T\n", age)

//
var address = "KIET Group Of Institutions, Ghaziabad"

fmt.Println("My Address is : ", address)

//
check := true

fmt.Println("Check is : ", check)

fmt.Println("Amount is : ", amount)

fmt.Println("Access token expiry is :", AccessTokenExpiry)


var dummy int
fmt.Println("Dummy value is :", dummy)

var isActive bool 
fmt.Println("Dummy value is :", isActive)

//Error- 

// n:= 5
// n:=10


// correct way-

n:= 5
fmt.Println("value of n is :", n)
n, s := 7, 10                       // atleast ek varibale new hona chahiye 

fmt.Println("value of n is :", n)
fmt.Println("value of s is :", s)

}