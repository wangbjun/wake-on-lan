#!/usr/bin/env bash
GOOS=linux GOARCH=mipsle go build -ldflags "-s -w"