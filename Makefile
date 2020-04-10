install:
	go mod vendor

test:
	go test --cover ./...

test-ci:
	GO111MODULE=off go get github.com/axw/gocov/gocov
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
	chmod +x ./cc-test-reporter
	chmod +x ./cover.sh
	./cc-test-reporter before-build
	go test -race -coverprofile c.out ./...
	./cc-test-reporter after-build

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run
