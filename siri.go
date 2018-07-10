package main

import (
	"fmt"
	"github.com/go-commandline-siri/airplane"
	"github.com/go-commandline-siri/tweet"

	"flag"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env.config")
}

func main() {

	tweetPtr := flag.Bool("tweet", false, "-tweet 'Tweet Goes here'")
	airplanePtr := flag.Bool("airplane", false, "-airplane '<on/off>'")

	flag.Parse()

	var ok bool

	if *tweetPtr {
		t := tweet.Tweet{flag.Args()[0]}
		ok = execute(t)
	} else if *airplanePtr {
		a := airplane.Airplane{Power: flag.Args()[0]}
		ok = execute(a)
	}

	if ok {
		fmt.Println("👍")
	} else {
		fmt.Println("👎")
	}
}

func execute(i Action) bool {
	return i.Execute()
}
