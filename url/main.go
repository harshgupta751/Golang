package main

import (
	"fmt"
	"net/url"
)

var myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymemtid=hfhghgj4552"

func main() {

	response, _ := url.Parse(myUrl)
	fmt.Println(response)
	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Port())
	fmt.Println(response.RawQuery)
	fmt.Println(response.Query())

	qParams:= response.Query()
	fmt.Printf("Type of qParams : %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, value := range qParams {
		fmt.Println(value)
	}

	partsOFUrl := &url.URL{
		Scheme: "https",
		Host: "harsh.dev",
		Path: "/developing",
		RawPath: "user=any",

	}
// fmt.Println(partsOFUrl)
	anotherUrl := partsOFUrl.String()
	fmt.Println(anotherUrl)


}
