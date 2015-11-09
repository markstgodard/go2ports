package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar")
}

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/foo", foo)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/bar", bar)

	go func() {
		fmt.Println("/foo listening on port 8081")
		http.ListenAndServe(":8081", mux1)
	}()

	fmt.Println("/bar listening on port 8082")
	http.ListenAndServe(":8082", mux2)
}
