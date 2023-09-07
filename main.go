package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":80", nil)
}

func root(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.String(), "/feed/") {
		handleFeed(w, r)
		return
	}

	redirect := ""
	switch r.URL.String() {

	case "/ping":
		w.Write([]byte("howdy"))

	case "/z00m":
		fallthrough
	case "/zoom":
		redirect = "https://zoom.us/j/9782268377"

	case "/highrock":
		redirect = "https://www.dropbox.com/sh/0zkf8wlizgwrl8o/AAA9L4txDf9cCGgTV9yQEX-7a"

	case "/rocks":
		redirect = "https://www.dropbox.com/sh/t5ab0zdszdv4pib/OQlqz9XmdP"

	case "/firmware":
		redirect = "https://notecard-firmware.s3.amazonaws.com/index.html"

	case "/fam1ly":
		fallthrough
	case "/family":
		redirect = "https://www.dropbox.com/sh/jc5vpxiihwsdvyu/IgWm75kT6U"

	default:
		fmt.Fprintf(w, "%s ???", r.URL)

	}

	if redirect != "" {
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	}

}
