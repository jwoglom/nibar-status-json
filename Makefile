main: FORCE
	go build -ldflags="-s -w" main.go
	upx main

FORCE: