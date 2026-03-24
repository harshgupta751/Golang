package main

import "fmt"

func main() {

	user := User{"Harsh", 19, "Ghaziabad", true}
	fmt.Println(user.Age);
	fmt.Println(user.getStatus())
	user.getAge()
	fmt.Println(user.Age);

}

type User struct {
	Name    string
	Age     int
	Address string
	Status  bool
}

func (u User) getStatus() bool {

return u.Status

}

func (u User) getAge(){
u.Age = 30
fmt.Println(u.Age);

}
