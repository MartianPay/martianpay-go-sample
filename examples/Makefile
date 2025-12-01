.PHONY: build run clean test help update

# Binary name
BINARY=examples

# Build the examples binary
build:
	@echo "Building examples..."
	@go build -o $(BINARY)
	@echo "✓ Build complete: $(BINARY)"

# Run the examples
run: build
	@echo "Running examples..."
	@./$(BINARY)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY)
	@echo "✓ Clean complete"

# Update MartianPay SDK to latest version
update:
	@echo "Updating MartianPay SDK to latest version..."
	@go get -u github.com/MartianPay/martianpay-go-sample
	@go mod tidy
	@echo "✓ SDK updated to latest version"

# Run go mod tidy
tidy:
	@echo "Running go mod tidy..."
	@go mod tidy
	@echo "✓ Dependencies updated"

# Format all Go files
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Formatting complete"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  make build   - Build the examples binary"
	@echo "  make run     - Build and run the examples"
	@echo "  make clean   - Remove build artifacts"
	@echo "  make update  - Update MartianPay SDK to latest version"
	@echo "  make tidy    - Run go mod tidy"
	@echo "  make fmt     - Format Go code"
	@echo "  make test    - Run tests"
	@echo "  make help    - Show this help message"

# Default target
.DEFAULT_GOAL := build
