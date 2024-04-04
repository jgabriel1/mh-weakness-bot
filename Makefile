.PHONY: build
build:
	@go build -o ./build/mh-weakness-bot ./cmd/bot/main.go

.PHONY: run
run:
	@go run ./cmd/bot/main.go