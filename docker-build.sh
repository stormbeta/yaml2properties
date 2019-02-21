#!/usr/bin/env bash

docker run --rm \
  -v "${PWD}:/build" \
  --workdir "/build" \
  -e GIT_COMMIT="${GIT_COMMIT:-$(git show-ref --head HEAD -s)}" \
  golang:1.11-alpine ./build.sh "$@"
