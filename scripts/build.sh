#!/usr/bin/env bash

GIT_COMMIT=$(git rev-parse HEAD)
LD_FLAGS="-X main.GitCommit=${GIT_COMMIT} $LD_FLAGS"

go build -ldflags="${LD_FLAGS}" -o ./bin/markdown-linter *.go
