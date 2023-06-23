FROM golang:latest as builder

ADD ../spmonitor-library /go/src/spmonitor-library
ADD ../service /go/src/service
WORKDIR /go/src/service

RUN go build -o service

EXPOSE 3333

CMD ["/go/src/service/service"]
