.PHONY:
.SILENT:

build:
  go build -o ./ cmd/bot/main.go

run: build
  ./

test:
  go test -v ./...