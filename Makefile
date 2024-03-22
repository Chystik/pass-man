SHELL=/bin/bash

.PHONY: dep
dep:
	go mod download
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -v -race ./...

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	/home/user/go/bin/golangci-lint run

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o=coverage.html
	rm coverage.out

.PHONY: dev-up
dev-up:
	docker-compose -f=docker-compose.dev.yml --env-file=.env up -d

.PHONY: dev-down
dev-down:
	docker-compose -f=docker-compose.dev.yml --env-file=.env down --rmi local

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    ./protobuf/pass_man.proto

buildVersion = 1.0.0
buildDate = $(shell date +'%Y/%m/%d %H:%M:%S')
buildCommit = $(shell git rev-parse HEAD)
.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -buildvcs=false -ldflags="-X main.buildVersion=1.0.0 -X main.buildDate=14.02.2024 -X main.buildCommit=assd" -o=bin/client-windows-amd64.exe ./cmd/client/
	GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-X main.buildVersion=1.0.0 -X main.buildDate=14.02.2024 -X main.buildCommit=assd" -o=bin/client-linux-amd64 ./cmd/client/
	GOOS=darwin GOARCH=amd64 go build -buildvcs=false -ldflags="-X main.buildVersion=1.0.0 -X main.buildDate=14.02.2024 -X main.buildCommit=assd" -o=bin/client-darwin-amd64 ./cmd/client/
	GOOS=darwin GOARCH=arm64 go build -buildvcs=false -ldflags="-X main.buildVersion=1.0.0 -X main.buildDate=14.02.2024 -X main.buildCommit=assd" -o=bin/client-darwin-arm64 ./cmd/client/
	GOOS=windows GOARCH=amd64 go build -buildvcs=false -o=bin/server-windows-amd64.exe ./cmd/server/
	GOOS=linux GOARCH=amd64 go build -buildvcs=false -o=bin/server-linux-amd64 ./cmd/server/
	GOOS=darwin GOARCH=amd64 go build -buildvcs=false -o=bin/server-darwin-amd64 ./cmd/server/
	GOOS=darwin GOARCH=arm64 go build -buildvcs=false -o=bin/server-darwin-arm64 ./cmd/server/
