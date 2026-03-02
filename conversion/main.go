package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

fmt.Println("Enter value")
	input, _ := reader.ReadString('\n')
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if(err!=nil){
		fmt.Println(err)
	}else{
		fmt.Println(rating)
		fmt.Println("Adding 1 to the input :", rating + 1 )
	}



}
