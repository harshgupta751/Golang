package main

import (
	"MongoAPI/router"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Server is running on Port 8020")
	http.ListenAndServe(":8020", router.Router())


}