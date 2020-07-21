#!/usr/bin/env sh

export GOFLAGS='-mod=vendor'

if command -v git >/dev/null 2>&1; then
  gitCommit="$(git rev-parse HEAD)"
else
  gitCommit="$GIT_COMMIT"
fi

case "${1:-build}" in
  build)
    go build -ldflags "-X main.commitVersion=${gitCommit}"
    ;;
  test)
    go test
    ;;
  *)
    go "$@"
esac

