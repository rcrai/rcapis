#!/usr/bin/env bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROTO_DIR="rcapis/"

echo "BASE_DIR=$BASE_DIR"

if [ -z $GOPATH ]; then
  echo "error: env: GOPATH not exists!!"
  exit 1
fi

SRC_DIR="$GOPATH/src"


cmd_check() {
    hash $1 2>/dev/null
}

cmd_required() {
    if $(cmd_check $1); then
        echo "CHECK: $1 OK"
    else
        echo "CHECK: $1 FAILED"
        echo "Note: "$2
        exit 1
    fi
}

cmd_required protoc "You can install it using 'brew install protobuf' in macOS, or search protobuf for other platform"
cmd_required protoc-gen-go "You can install it using 'go get -u github.com/golang/protobuf/protoc-gen-go'"
cmd_required protoc-gen-grpc-gateway "You can install it using 'go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway'"
#cmd_required protoc-gen-swagger "You can install it using 'go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger'"
cmd_required protoc-gen-go "You can install it using 'go get -u github.com/golang/protobuf/protoc-gen-go'"
#cmd_required protoc-gen-grpc_python "You can install it using 'go get -u github.com/golang/protobuf/protoc-gen-go'"


for file in $(find ${PROTO_DIR} -type f -name "*.pb.go" -o -name "*.pb.gw.go"); do
    echo "cleaning file: $file"
    rm -f ${file}
done

for file in $(find ${PROTO_DIR} -type f -name "*.proto"); do
    echo "processing: $file"
    protoc \
        --go_out=plugins=grpc:${SRC_DIR}\
        --grpc-gateway_out=logtostderr=true:${SRC_DIR} \
            ${file}
#        --swagger_out=logtostderr=true:${SRC_DIR} \
done



