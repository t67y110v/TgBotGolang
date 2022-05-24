.PHONY:
.SILENT:

build:
  go build -o ./cmd/bot/main.go

test:
  go test -v ./...