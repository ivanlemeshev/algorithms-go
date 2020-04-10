install:
	go mod vendor

test:
	go test --cover ./...

test-ci:
	go get github.com/axw/gocov/gocov
	go test -race -coverprofile=c.out ./...
	gocov convert c.out >> coverage.json

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run
