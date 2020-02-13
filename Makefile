.PHONY: run

run:
	@go run ./main.go

.PHONY: build

build:
	@go build -o ./build/server ./main.go
	@echo "[OK] Server was build!"
