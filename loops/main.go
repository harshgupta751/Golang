package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for i:=0; i<len(days); i++ {
		if(i==5){
			goto oct
		}

		fmt.Println(days[i])
	} 

	// i:=0
	// for i<len(days) {
	// 	fmt.Println(days[i])
	// 	i++;
	// }

	// for _, val:= range days {
	// 	fmt.Println(val)
	// }

oct:
		fmt.Println("Golang--")


}