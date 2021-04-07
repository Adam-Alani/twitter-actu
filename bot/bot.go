package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

func sendTweet( msg string, client *twitter.Client ) {
	tweet, _, _ := client.Statuses.Update(msg, nil)
	fmt.Println(tweet);
}

func fetchLatestTweet(client *twitter.Client, username string) string {
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: username, Count: 1}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)

	return tweets[0].Text
}



type Substitution struct {
	original string
	replacement string
}


var subDictionary = []Substitution {
	{"temoins", "mecs que je connais"},
	{"nouvelle etude", "post tumblr"},
	{"ecole", "hello"},
	{"macron", "manu"},
	{"eleves", "grosses merdes de la societé"},
	{"elections", "councours de restauration"},
	{"debat", "dance-off"},
	{"scientifique", "channing tatum et ses potes"},
	{"tension", "tension sexuelle"},
	{"americains", "obeses"},
	{"covid", "couilles vides "},
	{"minute", "annee"},
	{"annee ", "minute"},
	{"pharmacie", "bar"},
	{"confinement", "vacances"},
	{"vaccin", "puce electronique de Bill Gates"},
	{"vol", "emprunt"},
	{"suspect", "mec chelou"},
	{"electrique", "nucleaire"},
	{"voiture", "chat"},
	{"train", "chien"},
}



