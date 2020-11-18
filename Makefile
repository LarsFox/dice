default:
	@cat Makefile

run:
	go run cmd/main.go

build:
	go build -o dice cmd/main.go

linux:
	env GOOS=linux GOARCH=amd64 make build

mac:
	env GOOS=darwin GOARCH=amd64 make build

win:
	env GOOS=windows GOARCH=amd64 go build -o dice.exe cmd/main.go

lint:
	@golangci-lint run
