FROM golang:latest

RUN mkdir /go/src/app
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/app

WORKDIR /go/src/app

RUN dep ensure
RUN go build

CMD ["./app"]