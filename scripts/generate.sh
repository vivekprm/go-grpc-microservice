#/bin/bash

protoc --go_out=./pbgo --go-grpc_out=./pbgo proto/greet.proto 