export POSTGRES_USER=root
export POSTGRES_PASSWORD=pa55word
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export DATABASE_DRIVER="postgres"
export DATABASE_CONFIG="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/postgres?sslmode=disable"
go get -t
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
go get github.com/mattes/migrate
psql -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"
echo "CREATE EXTENSION IF NOT EXISTS pgcrypto" | psql -d postgres
migrate -url $DATABASE_CONFIG -path ./store/migrations up
go test ./... -cover -v
migrate -url $DATABASE_CONFIG -path ./store/migrations down
