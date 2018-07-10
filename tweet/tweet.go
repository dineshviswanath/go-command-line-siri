package tweet

import (
	"github.com/dghubble/oauth1"
	"os"

	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

// ErrInValidParam Throw this when invalid param passed to Action
var ErrInValidParam = errors.New("Invalid param passed")

// ErrUnableToExecute Throw this when unable to execute the Action
var ErrUnableToExecute = errors.New("Unable to execute")

// Tweet Holds Twitter.com interaction
type Tweet struct {
	Message string
}

// Validate Check if the Tweet param is valid
func (t Tweet) Validate() (error, bool) {
	if len(t.Message) < 5 {
		return ErrInValidParam, false
	} else {
		return nil, true
	}
}

//Execute Sends tweet
func (t Tweet) Execute() bool {

	err, _ := t.Validate()
	if err != nil {
		return false
	}

	isTweetSent := false

	consumerKey := os.Getenv("ENV_TWT_CONSUMER_KEY")
	consumerSecret := os.Getenv("ENV_TWT_CONSUMER_SECRET")
	tokenKey := os.Getenv("ENV_TWT_TOKEN_KEY")
	tokenSecret := os.Getenv("ENV_TWT_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(tokenKey, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Tweet client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err = client.Statuses.Update(t.Message, nil)
	if err != nil {
		fmt.Print(`Error: `)
		fmt.Print(err)
	} else {
		isTweetSent = true
	}
	return isTweetSent
}
