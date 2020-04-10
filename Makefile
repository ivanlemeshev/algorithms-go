install:
	go mod vendor

test:
	go test --cover ./...

test-ci:
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
	chmod +x ./cc-test-reporter
	./cc-test-reporter before-build
	go test -race -coverprofile c.out ./...
	sed -i -e 's/github.com\/ivanlemeshev\/algorithms-go\///g' c.out
	./cc-test-reporter after-build

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run
