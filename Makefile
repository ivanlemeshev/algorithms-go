install:
	go mod vendor

test:
	go test --cover ./...

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run
