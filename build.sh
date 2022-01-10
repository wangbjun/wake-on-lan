#!/bin/bash

GOARM=7 GOARCH=arm GOOS=linux go build -o autoWol -ldflags="-s -w"