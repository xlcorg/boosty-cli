
GO_PACKAGES=$(shell go list ./...)

.PHONY: run
run:
	go run cmd/boosty/main.go

.PHONY: build
build:
	go build -o bin/boosty cmd/boosty/main.go

.PHONY: docker-build
docker-build:
	docker build -t go-boosty-downloader .

.PHONY: test
test:
	go test ${GO_PACKAGES}