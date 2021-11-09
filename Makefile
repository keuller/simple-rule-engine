.PHONY: build

build:
	@go build -ldflags="-s- w" -o ./bin/princing-service main.go
	@echo "Binay generated"

run:
	@go run main.go

tests:
	@go test -v ./test/...
