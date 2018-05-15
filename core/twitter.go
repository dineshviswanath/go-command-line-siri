package core

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Twitter Holds Twitter.com interaction
type Twitter struct {
	TweetMessage string
}

// Validate Check if the Twitter param is valid
func (t Twitter) Validate() bool {
	return len(t.TweetMessage) > 0
}

//Execute Sends tweet
func (t Twitter) Execute() bool {
	isTweetSent := false

	consumerKey := os.Getenv("ENV_TWT_CONSUMER_KEY")
	consumerSecret := os.Getenv("ENV_TWT_CONSUMER_SECRET")
	tokenKey := os.Getenv("ENV_TWT_TOKEN_KEY")
	tokenSecret := os.Getenv("ENV_TWT_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(tokenKey, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update(t.TweetMessage, nil)
	if err != nil {
		fmt.Print(`Error: `)
		fmt.Print(err)
	} else {
		isTweetSent = true
	}
	return isTweetSent
}
