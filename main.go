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
	fmt.Fprintf(w, "http.Request\n")
	fmt.Fprintf(w, "%+v", r)
	fmt.Fprintf(w, "\nhttp.Request.URL\n")
	fmt.Fprintf(w, "%+v", r.URL)
	fmt.Fprintf(w, "\n")

	//	http.Redirect(w, req, "https://zoom.us/j/9782268377", http.StatusMovedTemporarily)
}
