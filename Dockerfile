FROM golang

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
#RUN psql -c "CREATE USER $POSTGRES_ENV_POSTGRES_USER WITH PASSWORD '$POSTGRES_ENV_POSTGRES_PASSWORD';"
#RUN echo "CREATE EXTENSION IF NOT EXISTS pgcrypto" | psql -d postgres
#RUN migrate postgres://$POSTGRES_ENV_POSTGRES_USER:$POSTGRES_ENV_POSTGRES_PASSWORD@POSTGRES_PORT_5432_TCP_ADDR:POSTGRES_PORT_5432_TCP_PORT/postgres?sslmode=disable -path ./store/migrations up

EXPOSE 8080

CMD ["go", "run", "./main.go"]
