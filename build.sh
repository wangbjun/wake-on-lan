#!/usr/bin/env bash
export GO111MODULE=on
GOOS=linux GOARCH=mipsle go build -o autoWol -ldflags "-s -w"