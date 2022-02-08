###
# Copyright 2022-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
### 
BINARY_NAME=analyticsEngine
OUTPUT_DIR=bin

build:
	go build -o ${OUTPUT_DIR} ./...

run:
	./bin/${BINARY_NAME}
test:
	go test -coverprofile=c.out  ./...

build_and_run: build run

clean:
	go clean
	rm ${OUTPUT_DIR}/*
