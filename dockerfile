FROM golang:1.9 as builder
WORKDIR /go/src/github.com/dineshviswanath/go-commandline-siri/
RUN go get -d -v github.com/dghubble/go-twitter/twitter
RUN go get -d -v github.com/dghubble/oauth1
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/dineshviswanath/go-commandline-siri/app .
CMD ["./app"]  