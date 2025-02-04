package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

var counter int

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Counter = ", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Available only GET method")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "Counter enlarged by 1")
	} else {
		fmt.Fprintln(w, "Available only POST method")
	}
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	http.HandleFunc("/post", PostHandler)

	http.HandleFunc("/get", GetHandler)

	http.ListenAndServe("localhost:8080", nil)
}
