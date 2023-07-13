BINARY_NAME=mav-cli

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: init-go
init-go:
	go get -v ./...

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/mav/main.go

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: all
all: clean fmt run