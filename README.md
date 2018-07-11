# Personal Command-line Assistant


## Install without using Docker (🤨)

Pre-requisite:
 - Golang installed
 - git
```

cd ~/go/src/github.com/

git clone git@github.com:dineshviswanath/go-command-line-siri.git

mv sample.env.config .env.config

# Replace following 
export ENV_TWT_CONSUMER_KEY="3yBdRC..."
export ENV_TWT_CONSUMER_SECRET="1AButlZjbaD..."
export ENV_TWT_TOKEN_KEY="18602..."
export ENV_TWT_TOKEN_SECRET="1mau..."

go get -v -d

go build siri.go
./siri -tweet "<UsuallyUselessTweetGoesHere>"
```


## Install with Docker (🤩)

Pre-requisite:
 - docker installed

```
docker build -t my-go-commandline-siri .
docker run --rm my-go-commandline-siri ./app -tweet "<UsuallyUselessTweetGoesHere>"
```


## ToDo
- [x] Send Tweets
- [x] Add tests in project
- [ ] Start DND
- [ ] Clean docker images
- [ ] Clean desktop
- [ ] Clean trash bin
- [ ] Hide Notifications
