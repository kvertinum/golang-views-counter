run:
	@go run ./cmd/server/main.go

test:
	@go test ./internal/app/...

build:
	@go build -o ./start-server ./cmd/server/main.go

compose-start:
	@docker compose up --build -d

compose-stop:
	@docker compose down -v --remove-orphans