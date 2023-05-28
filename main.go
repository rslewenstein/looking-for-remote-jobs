package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive...")
}

func handleRequests() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":5000", myRoute))
}

func main() {
	handleRequests()
}