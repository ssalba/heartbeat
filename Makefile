buildTime=$(shell date -u "+%Y%m%d%I%M")
revVersion=$(shell git rev-parse --short HEAD)

build:
	go build -ldflags "-X main.Version=$(revVersion).$(buildTime)"