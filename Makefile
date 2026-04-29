
build: main.go
	mkdir -p build
	go build -o build/dechar main.go

run: build
	./build/dechar

install: build
	mkdir -p ~/.local/bin
	cp build/dechar ~/.local/bin
