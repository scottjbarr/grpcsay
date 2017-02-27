.PHONY: protoc

# Apps to compile
APPS = grpcsay_server grpcsay_client grpcsay_http

# Command to use when building for arbitrary target
GO_BUILD = go build

build: protoc
	mkdir -p bin
	$(foreach app,$(APPS), $(GO_BUILD) -o bin/$(app) cmd/$(app)/*.go &&) exit 0

protoc:
	protoc -I grpcsay/ grpcsay/grpcsay.proto --go_out=plugins=grpc:grpcsay

clean:
	rm -rf bin

deps:
	go get -t ./...

run-http:
	BIND=:9000 bin/grpcsay_http

run-grpc-server:
	BIND=:50051 bin/grpcsay_server

run-grpc-client:
	ADDRESS=127.0.0.1:50051 bin/grpcsay_client "say over grpc"
