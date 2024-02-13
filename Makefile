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