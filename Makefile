.PHONY: build
build:
	@go build -o ./build/mh-weakness-bot ./cmd/bot/main.go

.PHONY: run
run:
	@go run ./cmd/bot/main.go

.PHONY: add_guild
add_guild:
	@go run ./cmd/add_guild/main.go -guildID=$(guildID)

.PHONY: setup
setup:
	@go install github.com/cosmtrek/air@latest

.PHONY: dev
dev:
	@air -c ./build/.air.toml