current_time = $(shell date +%c)
git_description = $(shell git describe --always --dirty --tags --long)

.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -ldflags="-s -X 'main.buildTime=${current_time}' -X 'main.version=${git_description}'" -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -X 'main.buildTime=${current_time}' -X 'main.version=${git_description}'" -o=./bin/linux_amd64/api ./cmd/api