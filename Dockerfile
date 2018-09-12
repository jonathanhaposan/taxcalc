FROM golang:1.11

RUN mkdir /go/src/github.com/jonathanhaposan/taxcalc
WORKDIR /go/src/github.com/jonathanhaposan/taxcalc
ADD . /go/src/github.com/jonathanhaposan/taxcalc

RUN go get -d -v ./...
RUN go build
RUN go install -v ./...