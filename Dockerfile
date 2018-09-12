FROM golang:1.11

RUN mkdir -p /go/src/github.com/jonathanhaposan/taxcalc/
WORKDIR /go/src/github.com/jonathanhaposan/taxcalc/
ADD . /go/src/github.com/jonathanhaposan/taxcalc/

RUN go get -d -v ./...
RUN go get gopkg.in/DATA-DOG/go-sqlmock.v1
RUN go test ./...
RUN go build
RUN go install -v ./...