package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://taskflow-todos.netlify.app"

func main() {

	response, err := http.Get(url)

		handleError(err)

fmt.Printf("Type of response is %T\n", response )
// fmt.Printf("Type of response is %T\n", response.Body )
// fmt.Println(response)
dataBytes, err := ioutil.ReadAll(response.Body)

handleError(err)

fmt.Println(string(dataBytes))


}

func handleError(err error){
if(err!=nil){
	panic(err)
}


}
