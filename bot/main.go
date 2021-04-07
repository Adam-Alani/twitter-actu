package main

import (
	"fmt"
	"log"
)


func main() {

	creds := Credentials{
		AccessToken: "3067410923-L2YSMfMZjhyvovPHKKpDQlO3wHxOjiJ7g4amL3e",
		AccessTokenSecret: "02pkTaSt1uacTzlFuOuUbreU0A8M2SwFvtpGNwEvzprfW",

		ConsumerKey: "54oAqtEdjnzcbLInK5UNHOWoE",
		ConsumerSecret: "vkDDwkxFuDk9UczyWEkisR4o8nvXwEGsFMN4EG21WRbyrhwRxZ",
	}


	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	latest := fetchLatestTweet(client, "unebrosseadam")
	fmt.Printf("USER TIMELINE:\n%+v\n", latest)


}
