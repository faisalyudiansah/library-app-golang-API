include .env

run :
	go run .

start-nodemon:
	nodemon --exec go run main.go --signal SIGTERM

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out && rm -f coverage.out

create-mock:
	mockery --all --case underscore --output ./mocks

db-generate:
	psql -d "dbname='${DBNAME_DB}' port='${PORT_DB}' user='${USER_DB}' password='${PASSWORD_DB}' host='${HOST_DB}'" -f databases/db.sql