package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating : ")

	rating, err := reader.ReadString('\n');

	fmt.Println("Thanks for rating, ", rating )
	fmt.Printf("Type is : %T", rating )
	fmt.Println();
	fmt.Println(err )
	fmt.Println(len(rating));
}
