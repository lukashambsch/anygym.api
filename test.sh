# shell script to set up env and run tests
export GOENV=test
export POSTGRES_ENV_POSTGRES_USER=root
export POSTGRES_ENV_POSTGRES_PASSWORD=pa55word
export POSTGRES_PORT_5432_TCP_ADDR=localhost
export POSTGRES_PORT_5432_TCP_PORT=5432
export DATABASE_DRIVER="postgres"
export DATABASE_CONFIG="postgres://$POSTGRES_ENV_POSTGRES_USER:$POSTGRES_ENV_POSTGRES_PASSWORD@localhost:5432/postgres?sslmode=disable"
go get -t
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
#go get github.com/mattes/migrate
psql -c "CREATE USER $POSTGRES_ENV_POSTGRES_USER WITH PASSWORD '$POSTGRES_ENV_POSTGRES_PASSWORD';"
echo "CREATE EXTENSION IF NOT EXISTS pgcrypto" | psql -d postgres
migrate -url $DATABASE_CONFIG -path ./store/migrations up
go test ./... -cover -v
migrate -url $DATABASE_CONFIG -path ./store/migrations down
export GOENV=local
