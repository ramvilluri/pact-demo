.PHONY: run-consumer unit

run-consumer:
	@go run consumer/client/cmd/main.go

unit:
	@echo "--- 🔨Running Unit tests "
	go test -count=1