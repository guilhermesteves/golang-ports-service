NAME=ps
VERSION=latest

docker-server:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -a -installsuffix cgo -o main .
	docker build -t $(NAME):$(VERSION) .
	rm main

docker-build:
	go mod download
	go mod vendor
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -a -installsuffix cgo -o main .
	docker build -t $(NAME):$(VERSION) .

build:
	go mod download
	go mod vendor
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -a -installsuffix cgo -o main .

install:
	go mod download

run:
	export HOST="localhost"; \
    export PORT="9090"; \
    go run cmd/main.go

test:
	go test -race ./...