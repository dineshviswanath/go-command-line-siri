FROM golang:1.9 as builder
WORKDIR $GOPATH/src/github.com/dineshviswanath/go-commandline-siri/
ADD . .
RUN go-wrapper download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app siri.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/dineshviswanath/go-commandline-siri/app .
CMD ["./app"]  