GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY=reducerApp

all: runb
build:  #build binary
	$(GOBUILD) -o $(BINARY) -v ./app
clean: #remove binary
	rm -f $(BINARY)
run: #run without creating a binary
	$(GOCMD) run app/main.go app/router.go app/ServiceContainer.go
runb: #run after creating a binary
	$(GOBUILD) -o $(BINARY) -v ./app
	./$(BINARY)


gen-proto:
    protoc -I server/api/ -I${GOPATH}/src --go_out=plugins=grpc:api server/api/api.proto

build-server:
    go build -i -v -o bin/server ${GOPATH}/src/grpc-db-poc/server/server

build-client:
    go build -i -v -o bin/client ${GOPATH}/src/grpc-db-poc/client-api

