#!/bin/bash

rm -f ./protocol/*.pb.go
protoc --gogo_out=./protocol --proto_path=./protocol ./protocol/*.proto
