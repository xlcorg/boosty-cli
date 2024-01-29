
GO_PACKAGES=$(shell go list ./...)

FLAGS="-X main.version=v0.0.3"

.PHONY: run
run:
	go run cmd/boosty/main.go

.PHONY: build
build:
	go build -v -o bin/boosty -ldflags $(FLAGS) cmd/boosty/boosty.go

.PHONY: docker-build
docker-build:
	docker build -t go-boosty-downloader .

.PHONY: test
test:
	go test ${GO_PACKAGES}

.PHONY: install
install: 
	go install -ldflags $(FLAGS) cmd/boosty/boosty.go
