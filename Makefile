export GO111MODULE=on

all: clean bld

bld: gin-test

gin-test:
	go build -o bin/grpc-server ./src/server
	go build -o bin/grpc-client ./src/client

clean:
	@rm -f bin/grpc-server
	@rm -f bin/grpc-client

cleanlog:
	@rm -f log/*log*