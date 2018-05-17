package main

import (
	"fmt"
	"os"


	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env.config")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usuage: siri <tweet> <message_to_tweet>`)
		return
	}

	var ok bool

	// TODO: Add command parser flag

	if os.Args[1] == "tweet" {
		t := Tweet{message:os.Args[2]}
		ok = execute(t)
	} else {
		fmt.Println(`Usuage: siri tweet <message_to_tweet>`)
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