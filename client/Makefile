build:
	go mod download
	go mod vendor
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -a -installsuffix cgo -o main .

install:
	go mod download

run:
	export HOST="localhost"; \
    export PORT="8080"; \
    export FILE="./ports.json"; \
    go run cmd/main.go

test:
	go test -race ./...

docker-build:
	go mod download
	go mod vendor
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -a -installsuffix cgo -o main .
	docker build -t $(NAME):$(VERSION) .