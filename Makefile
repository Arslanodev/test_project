run:
	go run ./cmd/api/main.go

build: 
	go build -o bin/blog-app ./cmd/api/main.go