FROM golang:1.8

ENV GOENV test

RUN ls $GOPATH/src

ADD . /go/src/app/

WORKDIR /go/src/app

RUN go env

RUN go get -u github.com/golang/dep/...
# RUN dep ensure -v
RUN go build

EXPOSE 8080

ENTRYPOINT /go/src/app/gym-all-over
