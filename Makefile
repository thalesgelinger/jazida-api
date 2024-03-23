build:
	@go build -o bin/jazida

run: build
	@./bin/jazida

test:
	@go test -v ./...


