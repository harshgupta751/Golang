package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)



func main() {

// PerformGetRequest()

PerformPostRequest()
}

func PerformGetRequest() {
	response, err := http.Get("http://localhost:5000/get")

	if(err!=nil){
		panic(err)
	}

	dataBytes, err := ioutil.ReadAll(response.Body)

		if(err!=nil){
		panic(err)
	}

		defer response.Body.Close()


	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

var responseString strings.Builder
     charCount, _ := responseString.Write(dataBytes)

	 fmt.Println(charCount)

	 fmt.Println(responseString.String())

}


func PerformPostRequest(){


content := strings.NewReader(`
{
"Name" : "Harsh Gupta",
"Language" : "Golang",
"dummy" : 12
}
`)

res, err :=	http.Post("http://localhost:5000/post", "application/json", content)

if err!=nil {
	panic(err)
}

dataBytes, _ := ioutil.ReadAll(res.Body)

defer res.Body.Close()

fmt.Println(string(dataBytes))


}