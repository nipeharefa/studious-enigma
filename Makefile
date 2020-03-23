CMD=GO111MODULE=on go
IMAGE_NAME=""

run:
		go run *.go

test:
		go test -v -race -covermode=atomic -coverprofile=coverage.coverprofile ./...

bench:
		$(GOCMD) test -bench=. -benchmem ./...

lint:
		golangci-lint run --concurrency 4 --print-issued-lines=false --exclude-use-default=false --enable=golint --enable=goimports  --enable=unconvert --enable=unparam

.PHONY: test bench

