build-docs:
	swag init -g src/api/main.go
dev:
	air
test:
	go test ./...
run:
	make build-docs && go run ./src/api/main.go