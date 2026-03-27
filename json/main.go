package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string    	`json:"coursename"`
	Price    int	    
	Platform string     `json:"platform"`
	Password string     `json:"-"`
	Tags     []string   `json:"tags,omitempty"`
}

func main() {
// EncodeJson()
decodeJson()



}


func EncodeJson(){
		content := []course{
		{"Harsh Gupta", 1200, "https://harsh.dev.com", "anyPassword", []string{"binary-search", "arrays", "NumberPattern"}},
		{"Mohan Singh", 2000, "https://temp.dev.com", "tempPassword", []string{"linear-search", "arrays", "NumberPattern"}},
		{"Raman", 1200, "https://raman.dev.com", "dummyPassword", []string{"binary-search", "arrays", "NumberPattern"}},
		{"Shyam", 1800, "https://shyam.dev.com", "any", nil},
	}

	jsonData, err := json.MarshalIndent(content, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}


func decodeJson(){

jsonData := []byte(`
[
        {
                "coursename": "Harsh Gupta",
                "Price": 1200,
                "platform": "https://harsh.dev.com",
                "tags": [
                        "binary-search",
                        "arrays",
                        "NumberPattern"
                ]
        },
        {
                "coursename": "Mohan Singh",
                "Price": 2000,
                "platform": "https://temp.dev.com",
                "tags": [
                        "linear-search",
                        "arrays",
                        "NumberPattern"
                ]
        },
        {
                "coursename": "Raman",
                "Price": 1200,
                "platform": "https://raman.dev.com",
                "tags": [
                        "binary-search",
                        "arrays",
                        "NumberPattern"
                ]
        },
        {
                "coursename": "Shyam",
                "Price": 1800,
                "platform": "https://shyam.dev.com"
        }
]
`)

check := json.Valid(jsonData)

if check {
	fmt.Println("Json is valid")
} else {
	fmt.Println("Json is Invalid")
}

var data []course

 json.Unmarshal(jsonData, &data)

	 fmt.Printf("%#v", data)
	fmt.Println()

	var MapData []map[string]interface{}

	json.Unmarshal(jsonData, &MapData)

fmt.Println(MapData)

	fmt.Println()

	for k, v := range MapData[0] {
		fmt.Printf("Key is: %v, Value is: %v, Type is: %T\n", k, v, v)
	}

}


