#!/bin/bash

protoc --go_out=./ --go-grpc_out=./ --proto_path=./ **/*.proto

cp -rf github.com/zhoujiawei1/protos/* .

rm -rf github.com

bazel build ...
