package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Tweet struct {
	Body string `json:"body"`
}

func getTweetsAction(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	tweets := []Tweet{}
	c := session.DB("tweeter").C("tweets")
	err = c.Find(bson.M{}).All(&tweets)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tweets)
}

func createTweetAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new tweet")
}

func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		fmt.Fprint(w, "not supported")
	case "GET":
		getTweetsAction(w, r)
	case "POST":
		createTweetAction(w, r)
	}
}

func main() {
	http.HandleFunc("/tweets/", TweetsHandler)
	http.ListenAndServe(":8080", nil)
}
