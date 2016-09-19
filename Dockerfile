FROM golang:latest

WORKDIR /go/src/github.com/lukashambsch/gym-all-over

ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV POSTGRES_USER root
ENV POSTGRES_PASSWORD pa55word
ENV DATABASE_DRIVER postgres
ENV DATABASE_CONFIG "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/postgres?sslmode=disable"

ADD . /go/src/github.com/lukashambsch/gym-all-over
RUN go get -t
RUN go get github.com/onsi/ginkgo
RUN go get github.com/onsi/gomega
RUN go get github.com/mattes/migrate
RUN migrate $DATABASE_CONFIG -path ./store/migrations up

EXPOSE 80

CMD ["go", "run", "main.go"]
