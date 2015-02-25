.PHONY: build fmt lint test vet deps

exe=ps-generator

default: build

build: vet fmt lint test
	go build -v -o $(GOPATH)/bin/$(exe) -i .

run: build
	$(GOPATH)/bin/$(exe)

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt .

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint .

test:
	go test -test.v .

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet .

deps:
	go get -u github.com/go-sql-driver/mysql
