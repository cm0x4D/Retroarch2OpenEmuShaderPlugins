all: build

build:
	@go run main.go -i ./shaders_slang -o ./plugins

clean:
	rm -rf ./plugins