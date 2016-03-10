package main

import (
	"fmt"
	"net/http"
)

func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "get list of tweets")
		return
	} else if r.Method == "POST" {
		fmt.Fprint(w, "create a new tweet")
		return
	} else if r.Method == "DELETE" {
		fmt.Fprint(w, "delete tweet")
		return
	}
	fmt.Fprint(w, "not supported")
}

func main() {
	http.HandleFunc("/tweets/", TweetsHandler)
	http.ListenAndServe(":8080", nil)
}
