package main

import (
	"MongoAPI/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on Port 8020")
	log.Fatal(http.ListenAndServe(":8020", router.Router()))

}