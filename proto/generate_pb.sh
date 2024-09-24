#!/bin/bash

# Find all .proto files in the proto directory and its subdirectories
find -name "*.proto" | while read -r proto_file; do
    # Get the directory of the proto file
    dir=$(dirname "$proto_file")

    # Run protoc to generate .pb.go files in the same directory as the .proto file
    protoc --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           "$proto_file"
done
