# gRPC Echo

A go gRPC server that echos back any message you supply it

# Requirements
- Go 1.12
- protoc compiler + gRPC go plugin (if generating protos; see https://grpc.io/docs/tutorials/basic/go/)

# Getting Started

## Running Server

```
make go-server
```

## Sending server a message

```
make go-client MESSAGE="hello"
```

## Running tests

```
make go-test
```
