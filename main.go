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

	switch r.URL.String() {

	case "/z00m":
		fallthrough
	case "/zoom":
		http.Redirect(w, r, "https://zoom.us/j/9782268377", http.StatusTemporaryRedirect)

	default:
		fmt.Fprintf(w, "%s ???", r.URL)

	}

}
