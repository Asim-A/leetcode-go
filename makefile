all:    ## clean, format, build and unit test
	make clean-all
	make gofmt
	make build
	make test

install:    ## build and install go application executable
	go install -v ./...

clean:  ## go clean
	go clean

clean-all:  ## remove all generated artifacts and clean all build artifacts
	go clean -i ./...

test: ## generate grpc code and run short tests
	go test -v ./... -short

