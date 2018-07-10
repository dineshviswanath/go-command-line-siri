package main

import (
	"fmt"
	"os"

	"github.com/go-commandline-siri/airplane"
	"github.com/go-commandline-siri/tweet"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env.config")
}

func main() {
	if len(os.Args) != 3 {
		// TODO: make usage coming from Flag parser
		fmt.Println(`Usage: siri tweet <message_to_tweet>`)
		fmt.Println(`Usage: siri airplane <on/off>`)
		return
	}

	var ok bool

	// TODO: Add command parser flag

	if os.Args[1] == "tweet" {
		t := tweet.Tweet{os.Args[2]}
		ok = execute(t)

	} else if os.Args[1] == "airplane" {
		a := airplane.Airplane{Power: os.Args[2]}
		ok = execute(a)
	} else {
		fmt.Println(`Usage: siri tweet <message_to_tweet>`)
		fmt.Println(`Usage: siri airplane <on/off>`)
		return
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
