export DATABASE_DRIVER="postgres"
export DATABASE_CONFIG="postgres://root:pa55word@localhost:5432/postgres?sslmode=disable"
go get -t
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
go get github.com/mattes/migrate
psql -c "CREATE USER root WITH PASSWORD 'pa55word';"
migrate -url $DATABASE_CONFIG -path ./store/migrations up
go test ./handlers
migrate -url $DATABASE_CONFIG -path ./store/migrations down
migrate -url $DATABASE_CONFIG -path ./store/migrations up
go test ./store/datastore
migrate -url $DATABASE_CONFIG -path ./store/migrations down
