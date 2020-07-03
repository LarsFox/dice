default:
	@cat Makefile

linux:
	env GOOS=linux GOARCH=amd64 go build

mac:
	env GOOS=darwin GOARCH=amd64 go build

win:
	env GOOS=windows GOARCH=amd64 go build
