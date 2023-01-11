#!/bin/bash

mkdir -p ../generated/proto
protoc --go_out=../generated/proto --go_opt=paths=source_relative --go-grpc_out=../generated/proto --go-grpc_opt=paths=source_relative ./*.proto
sed -i "" -e "s/,omitempty//g" ../generated/proto/*.go