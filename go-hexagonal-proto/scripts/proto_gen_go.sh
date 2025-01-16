#!/usr/bin/env bash

# Licensed to the LF AI & Data foundation under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SCRIPTS_DIR=$(dirname "$0")
PROTO_DIR=$SCRIPTS_DIR/../proto/
PROGRAM=$(basename "$0")
GOPATH=$(go env GOPATH)
GOOGLE_PROTO_DIR=$SCRIPTS_DIR/../cmake-build/protobuf/protobuf-src/src/

if [ -z $GOPATH ]; then
    printf "Error: the environment variable GOPATH is not set, please set it before running %s\n" $PROGRAM > /dev/stderr
    exit 1
fi

case ":$PATH:" in
    *":$GOPATH/bin:"*) ;;
    *) export PATH="$GOPATH/bin:$PATH";;
esac

echo "using protoc-gen-go: $(which protoc-gen-go)"

pushd ${PROTO_DIR}

GO_API="go-api"

find ../${GO_API}/${module} -type f ! \( -iname 'go.mod' -o -iname 'go.sum' \) -exec rm -f {} +

PROTO=$(find ./${module} -maxdepth 10 -type f -path '*.proto' | sort)

protoc -I . \
    --validate_out="lang=go:../${GO_API}" \
    --validate_opt paths=source_relative \
    --go_out=../${GO_API} --go_opt=paths=source_relative \
    --go-grpc_out=../${GO_API} --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=../${GO_API} \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    $PROTO

popd