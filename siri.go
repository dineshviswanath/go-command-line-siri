package main

import (
	"fmt"
	"os"

	"github.com/go-commandline-siri/core/twitter"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env.config")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`Usuage: siri <message_to_tweet>`)
		return
	}

	msg := os.Args[1]
	// TODO: Add command parser flag

	t := twitter.Twitter{TweetMessage: msg}

	if ok := tweet(msg); ok {
		fmt.Println("Tweet sent :)")
	} else {
		fmt.Println("Tweet could not sent :(")
	}
}
