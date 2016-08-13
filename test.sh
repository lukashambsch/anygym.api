export GO_DB_URL="postgres://postgres:@localhost:5432/test_gym_all_over?sslmode=disable"
migrate -url $GO_DB_URL -path ./store/migrations up
go test ./handlers
migrate -url $GO_DB_URL -path ./store/migrations down
migrate -url $GO_DB_URL -path ./store/migrations up
go test ./store/datastore
migrate -url $GO_DB_URL -path ./store/migrations down
