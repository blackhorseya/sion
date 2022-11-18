#!/usr/bin/env bash

set -e
echo "" >coverage.txt

for d in $(go list ./... | grep -v pb | grep -v docs | grep -v mock); do
  go test -race -coverprofile=profile.out -covermode=atomic --tags=unit "$d"
  if [ -f profile.out ]; then
    cat profile.out | grep -v "mock" >>coverage.txt
    rm profile.out
  fi
done
