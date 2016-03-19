package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
	"log"
	"flag"
	"time"
)

type Tweet struct {
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Config struct {
	MongoUrl string `json:"mongo_url"`
	MongoDb string `json:"mongo_db"`
}

var mongo mgo.Session
var mongodb mgo.Database
var config Config

func getTweetsAction(w http.ResponseWriter, r *http.Request) {
	tweets := []Tweet{}
	c := mongodb.C("tweets")
	err := c.Find(nil).Sort("-created_at").All(&tweets)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tweets)
}

func createTweetAction(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet
	json.NewDecoder(r.Body).Decode(&tweet)

	c := mongodb.C("tweets")
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
	host := flag.String("host", "127.0.0.1", "Host")
	port := flag.String("port", "8080", "Port")
	flag.Parse()

	loadConfig("config."+*env+".json")

	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	mongo = *session
	mongodb = *session.DB(config.MongoDb)
	defer session.Close()

	http.HandleFunc("/tweets/", TweetsHandler)

	log.Fatal(http.ListenAndServe(*host+":"+*port, nil))
}
