
GO_PACKAGES=$(shell go list ./...)

.PHONY: run
run:
	go run cmd/boosty/main.go

.PHONY: build
build:
	go build -v -o bin/boosty -ldflags "-X main.version=v0.0.2" cmd/boosty/boosty.go

.PHONY: docker-build
docker-build:
	docker build -t go-boosty-downloader .

.PHONY: test
test:
	go test ${GO_PACKAGES}

.PHONY: build
install:
	go install cmd/boosty/boosty.go
