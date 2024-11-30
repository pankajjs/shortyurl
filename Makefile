dev:
	@air

build:
	@go build -o build/shortyurl main.go

run:build
	@./build/shortyurl