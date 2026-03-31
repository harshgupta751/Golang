package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("Get")

	log.Fatal(http.ListenAndServe(":8020", r))

}

func greeter() {
	fmt.Println("Mod in Golang")

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Response from Home API<h1/>"))

}
