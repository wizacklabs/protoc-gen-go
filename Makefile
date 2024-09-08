SOURCES=$(shell find . -name "*.go")
VERSION=$(shell git describe --tags --long --abbrev=8)
RC?=0

protoc-gen-go: $(SOURCES)
	go build -ldflags '-s -w -X main.version=$(VERSION) -X main.rc=$(RC)' -o $@ .


.PHONY: clean
clean:
	@rm -f protoc-gen-go