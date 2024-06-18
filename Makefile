include .env.example

build:
	@go build -o build/package/go-api-clean-arch cmd/main.go
run:
	@PORT=$(PORT) \
	 VERSION=$(VERSION) \
	./build/package/go-api-clean-arch
test:
	@go test -v ./...