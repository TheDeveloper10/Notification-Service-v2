get_dependencies:
	go get -v ./...

test:
	go test -v ./...

test_race:
	go test -v -race ./...

build:
	go build -o bin/main main.go

run:
	go run main.go