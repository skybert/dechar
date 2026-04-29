VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR := build


build: main.go
	mkdir -p $(BUILD_DIR)
	go build \
		-ldflags "-X main.Version=$(VERSION)" \
		-o $(BUILD_DIR)/dechar \
		main.go

fmt: main.go
	gofmt -w -s .

install: build
	mkdir -p ~/.local/bin
	cp $(BUILD_DIR)/dechar ~/.local/bin

clean:
	rm -rf $(BUILD_DIR)

vuln:
	govulncheck ./...

.PHONY: fmt clean
