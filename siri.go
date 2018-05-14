package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`Usuage: siri <message_to_tweet>`)
		return
	}

	msg := os.Args[1]

	if ok := tweet(msg); ok {
		fmt.Println("Tweet sent :)")
	} else {
		fmt.Println("Tweet could not sent :(")
	}
}

func tweet(msg string) bool {

	isTweetSent := false

	consumerKey := os.Getenv("ENV_TWT_CONSUMER_KEY")
	consumerSecret := os.Getenv("ENV_TWT_CONSUMER_SECRET")
	tokenKey := os.Getenv("ENV_TWT_TOKEN_KEY")
	tokenSecret := os.Getenv("ENV_TWT_TOKEN_SECRET")

	printSmt := []string{consumerKey, consumerSecret, tokenKey, tokenSecret}
	fmt.Println(printSmt)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	fmt.Println(config)
	token := oauth1.NewToken(tokenKey, tokenSecret)
	fmt.Println(token)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update(msg, nil)
	if err != nil {
		fmt.Print(`Error: `)
		fmt.Print(err)
	} else {
		isTweetSent = true
	}

	return isTweetSent
}
