# grpc-practice

## Requirements

OS Version : MacOS

Go Version : 1.11 (with GoModule)

## Get started

1. If your os version is not macOS, you can download and replace the bin/protoc from
https://github.com/google/protobuf/releases.
2. If your go version is not 1.11, you can install gvm to checkout go to different version.
3. Follow the steps below
```go
$ export GO111MODULE

$ ./bin/protoc --go_out=plugins=grpc:. ./src/proto/*.proto

$ make
```

4. Exeucte Client-Server Program by following steps
```go
// Execute server by
$ ./bin/grpc-server

// Execute client by
$ ./bin/grpc-client
```