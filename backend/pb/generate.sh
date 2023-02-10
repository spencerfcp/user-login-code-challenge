#!/bin/bash
set -euo pipefail

ROOT="$(git rev-parse --show-toplevel)"
PB_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"
SCOIR_FRONTEND="${ROOT}/frontend"
function sedi() {
    # sed uses a different arg to omit backups in mac vs linux.  Use this function to work on both.
    if [[ "${OSTYPE}" == "darwin"* ]]; then
        sed -i '' "$@"
    else
        sed -i'' "$@"
    fi
}

GOPATH=$(go env GOPATH)

compileGOPb() {
    GO_FILENAME="${1}"
    SOURCE_FILENAME="${2}"

    echo "Compiling ${SOURCE_FILENAME} for Scoir API"
    protoc \
        --plugin="${GOPATH}/bin/protoc-gen-go" \
        --go_out=. "./${SOURCE_FILENAME}"

    sedi -e "s/protoc[ ]*v3.[0-9]*.[0-9]*/protoc        v3.x.x/g" "$GO_FILENAME"
    printf '%s\n\n%s\n' "$(cat "${GO_FILENAME}")" "// md5-hash $(md5 -q "./${SOURCE_FILENAME}")" >"${GO_FILENAME}"
}

compileAppPb() {
    MODULE="${1}"
    LC_MODULE=$(echo "${MODULE}" | awk '{ print tolower($1) }')
    echo "Compiling ${MODULE} for Scoir App"
    protoc \
        --plugin="${ROOT}/frontend/node_modules/ts-proto/protoc-gen-ts_proto" \
        --ts_proto_out="${SCOIR_FRONTEND}/pb" \
        --ts_proto_opt=outputEncodeMethods=false \
        --ts_proto_opt=snakeToCamel=false \
        "./api.proto"
}

# Unable to install the vendor version of this plugin due to the internal dependency
# Installing globally but checking version
CURRENT_VERSION=""
if hash "${GOPATH}/bin/protoc-gen-go" 2>/dev/null; then
    CURRENT_VERSION=$("${GOPATH}/bin/protoc-gen-go" --version)
fi

if [[ "${CURRENT_VERSION}" != "protoc-gen-go v1.28.0" ]]; then
    echo "Installed new protoc-gen-go version"
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
fi

pushd "${PB_DIR}"

compileGOPb api.pb.go api.proto

compileAppPb api

echo successfully completed
