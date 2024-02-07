#!/bin/bash

# clean output directory and create output directory
# rm -rf out
# mkdir -p out

# sandbox
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    sandbox/sandbox.proto

python -m grpc_tools.protoc -I./sandbox --python_out=./sandbox --pyi_out=./sandbox --grpc_python_out=./sandbox --proto_path sandbox/sandbox.proto sandbox/sandbox.proto
