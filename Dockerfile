FROM golang:1.8

ENV GOENV test

ADD . /go/src/github.com/lukashambsch/anygym.api/

WORKDIR /go/src/github.com/lukashambsch/anygym.api

RUN go get -u github.com/golang/dep/...
# RUN dep ensure -v
RUN go build

EXPOSE 8080

CMD ["./anygym.api"]
