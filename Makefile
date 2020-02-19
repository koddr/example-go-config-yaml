.PHONY: all

all: test build run

.PHONY: run

run:
	@go run ./main.go

.PHONY: build

build:
	@go build -o ./build/server ./main.go
	@echo "[OK] Server was build!"

.PHONY: test

test:
	@go test -v -coverprofile=cover.out ./...

.PHONY: show_coverage

show_coverage:
	@go tool cover -html=cover.out
