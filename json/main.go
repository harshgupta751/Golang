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

	content := []course{
		{"Harsh Gupta", 1200, "https://harsh.dev.com", "anyPassword", []string{"binary-search", "arrays", "NumberPattern"}},
		{"Mohan Singh", 2000, "https://temp.dev.com", "tempPassword", []string{"linear-search", "arrays", "NumberPattern"}},
		{"Raman", 1200, "https://harsh.dev.com", "anyPassword", []string{"binary-search", "arrays", "NumberPattern"}},
		{"Raman", 1200, "https://harsh.dev.com", "anyPassword", nil},
	}

	jsonData, err := json.MarshalIndent(content, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)

}
