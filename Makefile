build:
	@go build -o ./bin/vinom-common ./main.go

test:
	go test -v ./...

run: build
	@./bin/vinom-common
