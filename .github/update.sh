#!/bin/bash

echo "> installing protoc"
export GO111MODULE=on
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u google.golang.org/grpc
go install github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

echo "> removing old directories"
rm -rf discord
rm -rf gateway

echo "> cloning fresh protos..."
rm -rf proto
git clone https://github.com/captcha-bot/protobufs.git proto

echo "> generating code"
protoc -Iproto --go_out=. \
    --go-grpc_out=. \
    $(find ./ -type f -name "*.proto" | grep -v google)

echo "> re-organizing directories..."
mv ./github.com/captcha-bot/protobuf-go/* ./

echo "> go mod tidy"
go mod tidy
