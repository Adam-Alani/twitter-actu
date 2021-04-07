package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Credentials struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func getCredentials() Credentials {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Credentials{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret:  os.Getenv("ACCESS_SECRET"),

		ConsumerKey:  os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:  os.Getenv("CONSUMER_SECRET"),
	}
}


func getClient(creds *Credentials) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}


	return client, nil

}


