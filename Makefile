.phony: fmt all vet dir lint tst cov tstv

build_lib = build

all: fmt

dir:
	mkdir -p $(build_lib)

fmt:
	go fmt ./pkg/...

vet: all
	go vet ./pkg/...

lint: fmt
	golangci-lint run -v --timeout 5m

tst: fmt
	go test -mod=mod -coverprofile=coverage.out ./pkg/...

cov: tst
	go tool cover -html=coverage.out

tstv: fmt
	go test -v -mod=mod -coverprofile=coverage.out ./pkg/...
