SOURCES=$(shell find . -name "*.go")

protoc-gen-go: $(SOURCES)
	go build -ldflags '-s -w' -o $@ .