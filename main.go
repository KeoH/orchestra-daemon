package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Image struct {
	Name string
}

type Data struct {
	Images []Image
}

func serializeJSON(data Data, response http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func indexView(response http.ResponseWriter, request *http.Request) {
	data := Data{[]Image{Image{"docker"}, Image{"mongodb"}}}
	serializeJSON(data, response, request)
}

func main() {
	http.HandleFunc("/", indexView)
	fmt.Println("Serving on port 12001")
	http.ListenAndServe(":12001", nil)
}
