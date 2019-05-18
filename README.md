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

## Generating new TLS CA and Keys

Install certstrap found here https://github.com/square/certstrap

Delete the folder `tls` and remove existing certs and CA files

```
certstrap --depot-path "tls" init --common-name "Echo CA"
certstrap --depot-path "tls" request-cert --common-name Client -ip "127.0.0.1"
certstrap --depot-path "tls" sign Client --CA "Echo CA"
```