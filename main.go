package main

import (
	"fmt"
	"net/http"
)

func indexView(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello")
}

func main() {
	http.HandleFunc("/", indexView)
	fmt.Println("Serving on port 12001")
	http.ListenAndServe(":12001", nil)
}
