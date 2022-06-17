#!/bin/bash

protoc -I . --gogo_out=plugins=grpc:../internal/pb ./*.proto
