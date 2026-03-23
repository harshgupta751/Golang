package main

import "fmt"

func main() {

	u := User{"Harsh", 19, "Ghaziabad", true}
	fmt.Println(u)
	fmt.Println(u.Name)
	fmt.Println(u.Status)
	fmt.Printf("User %v\n", u )
	fmt.Printf("User %+v\n", u )


}

type User struct {
	Name    string
	Age     int
	Address string
	Status  bool
}
