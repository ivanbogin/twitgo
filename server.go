package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"log"
	"flag"
)

type Tweet struct {
	Body string `json:"body"`
}

type Config struct {
	MongoUrl string `json:"mongo_url"`
}

var mongo mgo.Session
var config Config

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

func loadConfig(filename string) {
	configFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(configFile).Decode(&config)
	configFile.Close()
}

func main() {
	env := flag.String("env", "production", "Environment")
	flag.Parse()

	loadConfig("config."+*env+".json")

	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	mongo = *session
	defer session.Close()

	http.HandleFunc("/tweets/", TweetsHandler)

	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}
