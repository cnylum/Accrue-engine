.PHONY: build test lint proto clean docker sdk-build sdk-clean

SERVICES := gateway adapter portfolio ledger

## Build all services
build:
	@for svc in $(SERVICES); do \
		echo "=== Building $$svc ==="; \
		cd services/$$svc && go build -o ../../bin/$$svc ./cmd/... && cd ../..; \
	done

## Run all tests
test:
	@for svc in $(SERVICES); do \
		echo "=== Testing $$svc ==="; \
		cd services/$$svc && go test ./... -race -count=1 && cd ../..; \
	done
	@echo "=== Testing SDK ==="
	@cd sdk && go test ./... -race -count=1

## Run linter
lint:
	golangci-lint run ./...

## Generate protobuf code
proto:
	buf generate
	@echo "Proto code generated in gen/go/"

## Lint protobuf definitions
proto-lint:
	buf lint

## Build SDK shared library
sdk-build: sdk-build-shared sdk-build-wasm

sdk-build-shared:
	@echo "=== Building SDK shared library ==="
	cd sdk && CGO_ENABLED=1 go build -buildmode=c-shared -o ../bin/libaccrue.so ./sharedlib/

sdk-build-wasm:
	@echo "=== Building SDK WASM ==="
	cd sdk && GOOS=js GOARCH=wasm go build -o ../bin/accrue.wasm ./wasm/

sdk-clean:
	rm -f bin/libaccrue.so bin/libaccrue.h bin/accrue.wasm

## Build Docker images
docker:
	@for svc in $(SERVICES); do \
		echo "=== Docker build $$svc ==="; \
		docker build -t accrue-$$svc:dev services/$$svc; \
	done

## Start all services
up:
	docker-compose up

## Start all services (detached)
up-d:
	docker-compose up -d

## Stop all services
down:
	docker-compose down

## Clean build artifacts
clean: sdk-clean
	rm -rf bin/
