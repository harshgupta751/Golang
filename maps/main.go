package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("Languages are:", languages);

	delete(languages, "rb");

	fmt.Println("Remaining Languages are:", languages);

	for _, value := range languages {
		fmt.Printf("Value: %v\n", value)
	}

}
