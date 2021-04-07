package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"time"
)

func main() {

	username := "BFMTV"


	// Initialize creds and client
	creds := getCredentials()
	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// Initialize first tweet and the substitution library
	old := fetchLatestTweet(client, username)
	dict := parseSub("./substitution_dictionary")


	// Fetches new tweet every minute
	for range time.NewTicker(time.Minute).C {
		latest := mainLoop(old, username, client , dict)
		old = latest
	}

}


func mainLoop(old , username string, client *twitter.Client, dict []Substitution) string {

	latest := fetchLatestTweet(client, username)
	if latest != old {
		modifiedTweet, modifiable := modifyTweet(latest, dict)
		article, postable := isArticle(latest)
		if  postable && modifiable {
			tweetToSend := modifiedTweet + "\n" + article
			sendTweet(tweetToSend, client)
		}
	}
	return latest
}
