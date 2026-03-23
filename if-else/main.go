package main

import "fmt"

func main() {

	loginCount := 26

		var result string
	if loginCount < 10 {
		result= "LoginCount is less than 10"
	} else if loginCount>10 {
			result= "LoginCount is more than 10"
	} else {
			result= "LoginCount is equal to 10"
	}
fmt.Println(result)

	if 9%2==0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	
	if num:=53; num>50 {
		fmt.Println("Num is greater than 50")
	} else {
		fmt.Println("Num is less than or equal to 50")
	}

}