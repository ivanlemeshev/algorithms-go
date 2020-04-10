install:
	go mod vendor

test:
	go test --cover ./...

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

test-ci:
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
	chmod +x ./cc-test-reporter
	chmod +x ./cover.sh
	go get github.com/axw/gocov/gocov
	./cover.sh
	./cc-test-reporter after-build --exit-code $TRAVIS_TEST_RESULT
