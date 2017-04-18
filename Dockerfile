FROM golang:1.8

ENV GOENV test

ADD . /go/src/github.com/lukashambsch/gym-all-over/

WORKDIR /go/src/github.com/lukashambsch/gym-all-over

RUN go get -u github.com/golang/dep/...
# RUN dep ensure -v
RUN go build

EXPOSE 8080

CMD ["./gym-all-over"]
