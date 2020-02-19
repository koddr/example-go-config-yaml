.PHONY: all

all: test build run

.PHONY: run

run:
	@go run ./...

.PHONY: build

build:
	@go build -o ./build/server ./...
	@echo "[OK] Server was build!"

.PHONY: test

test:
	@go test -v -coverprofile=cover.out ./...
	@echo "[OK] Test and coverage file was created!"

.PHONY: show_coverage

show_coverage:
	@go tool cover -html=cover.out
	@echo "[OK] Coverage file opened!"
