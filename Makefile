build:
	@go build -o bin/jazida cmd/jazida/main.go

run: build
	@./bin/jazida

test:
	@go test -v ./...


