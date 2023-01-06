all: test vet build
test:
	go test ./...
vet:
	go vet ./...
build:
	go build -o ./bin/api ./src/cmd/api
clean:
	rm -r ./bin