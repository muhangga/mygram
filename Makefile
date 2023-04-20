run:
	nodemon --exec go run cmd/main.go --ext go

build:
	go build -o bin/main cmd/main.go