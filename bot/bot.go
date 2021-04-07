package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"mvdan.cc/xurls/v2"
	"net/http"
	"strings"
	"unicode"
)

func sendTweet( msg string, client *twitter.Client ) {

	tweet, _, _ := client.Statuses.Update(msg, nil)
	fmt.Println(tweet);
}

func fetchLatestTweet(client *twitter.Client, username string) string {

	userTimelineParams := &twitter.UserTimelineParams{
		ScreenName: username,
		Count: 15,
		TweetMode: "extended",
		IncludeRetweets: &[]bool{false}[0],
		ExcludeReplies: &[]bool{true}[0],
	}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)

	i := 0
	for i < len(tweets) {
		if !tweets[i].Retweeted {
			if !tweets[i].Truncated { //IDK WHY ITS WHEN ITS NOT TRUNCATED THAT I GET FULL TEXT
				return tweets[i].FullText
			} else {
				return tweets[i].Text
			}
		}
		i++
	}
	return tweets[0].FullText
}


func isArticle(tweet string) (string, bool) {

	urls := xurls.Relaxed().FindAllString(tweet ,-1)
	for i := range urls {

		// Open the redirect URL to find the true irl, and replace the redirected one with real
		resp , _ := http.Get(urls[i])
		urls[i] = resp.Request.URL.String()

		// This means that one of the links is an article
		if !strings.Contains(urls[i], "/twitter.com/") {
			return urls[i], true
		}
	}

	return "", false
}


func modifyTweet(tweet string, dict []Substitution) (string, bool) {

	// Removes all accents and weird characters from text, é -> e
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	newTweet,_,_ := transform.String(t,strings.ToLower(tweet))

	for _, substitution := range dict {
		newTweet = strings.ReplaceAll(newTweet, substitution.original, substitution.replacement)
	}

	original, _ ,_ := transform.String(t,strings.ToLower(tweet))


	if original != newTweet { return newTweet, true }
	return "", false
}






