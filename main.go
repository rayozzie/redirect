package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":80", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
