clean:
	@echo "Cleaning up..."
	@rm -rf build/

build: clean
	@echo "Building project..."
	@go build -o build/api

run: build
	@echo "Running project..."
	@./build/api

test: 
	@go test -v ./tests/...