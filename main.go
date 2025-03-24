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

	case "/archive":
		redirect = "https://blues-event-archive.s3.amazonaws.com/index.html"

	case "/fam1ly":
		fallthrough
	case "/family":
		redirect = "https://www.dropbox.com/sh/jc5vpxiihwsdvyu/IgWm75kT6U"

	case "/temp":
		fmt.Fprintf(w, `PUT authorization {"slot": 0, "signed_digest": "424c700200000001c63dc3ed1648666dc3027c82cf003063972b03bf8d382d154c021182674a4b21031bec77bf874c41694dcc2b270518798e70332997e844909d700e25e3afccc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000504e4631413031345100c800040200000000c9000438539de100ca000400000000ffff00000000000000000000000000"}`)

	default:
		fmt.Fprintf(w, "%s ???", r.URL)

	}

	if redirect != "" {
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	}

}
