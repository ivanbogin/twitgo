package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tweet struct {
	Body string `json:"body"`
}

var tweets = []Tweet{Tweet{Body: "First tweet"}, Tweet{Body: "Second tweet"}, Tweet{Body: "Third tweet"}}

func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	if r.Method == "GET" {
		enc.Encode(&tweets)
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
