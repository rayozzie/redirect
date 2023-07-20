package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handleFeed(w http.ResponseWriter, r *http.Request) {

	filename := strings.TrimPrefix(r.URL.String(), "/feed/")
	fmt.Printf("get feed for %s\n", filename)
	if filename == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	auth := false
	u, p, ok := r.BasicAuth()
	if ok {
		// curl -vv 'http://abc:123@ozz.ie/feed/test'
		// {"req":"web.get","route":"feed"} with route being that same thing
		var (
			username = "abc"
			password = "123"
		)
		if u != username {
			fmt.Printf("Username provided is correct: %s\n", u)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if p != password {
			fmt.Printf("Password provided is correct: %s\n", u)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		auth = true
	}

	contents, err := os.ReadFile("feeds/filename")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Set cache paramter (max-age is in units of seconds)
	_ = auth
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(contents)

}
