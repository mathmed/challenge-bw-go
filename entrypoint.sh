#!/bin/bash
cd /usr/local/go/src/challenge-bw-go && /usr/local/go/bin/go mod tidy
exec /usr/local/go/bin/go run /usr/local/go/src/challenge-bw-go/src/main.go