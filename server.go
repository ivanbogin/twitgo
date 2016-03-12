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

var mongo mgo.Session

func getTweetsAction(w http.ResponseWriter, r *http.Request) {
	tweets := []Tweet{}
	c := mongo.DB("tweeter").C("tweets")
	err := c.Find(bson.M{}).All(&tweets)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tweets)
}

func createTweetAction(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet
	json.NewDecoder(r.Body).Decode(&tweet)

	c := mongo.DB("tweeter").C("tweets")
	err := c.Insert(&tweet)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
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
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	mongo = *session
	defer session.Close()

	http.HandleFunc("/tweets/", TweetsHandler)
	http.ListenAndServe(":8080", nil)
}
