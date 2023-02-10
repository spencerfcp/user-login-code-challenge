#!/bin/bash
set -euo pipefail

xcode-select --install || true

brew update
brew upgrade
brew install fswatch

brew install go
brew install wget
brew install node

# npm install -g prettier

brew install pnpm

# protobuf
brew install protobuf
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh |
    sh -s -- -b "$(go env GOPATH)/bin" v1.50.0

brew install direnv
brew install docker --cask

echo setup dev machine completed successfully
