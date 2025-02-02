package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	greeting := "Hello my friend"
	fmt.Fprintf(writer, "%s", greeting)
}
