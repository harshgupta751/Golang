package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Create("./DemoFile.txt")
	handleError(err)
	fmt.Println(file)
	content := "This content is to be written in file"

	len , err := io.WriteString(file, content)
		handleError(err)
		fmt.Println("Length is :", len)
	defer file.Close()
	readFile("./DemoFile.txt")
}

func handleError(err error){

if(err!=nil){
	panic(err)
}

}

func readFile(fileName string){

data, err := ioutil.ReadFile(fileName)

handleError(err)
fmt.Println(string(data))


}