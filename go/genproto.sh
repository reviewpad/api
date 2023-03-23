#!/bin/bash -eu
# Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
# Use of this source code is governed by a license that can be
# found in the LICENSE file.

protodir=../pb

rm -rf ./entities
rm -rf ./services
rm -rf ./codehost

protoc --go_out=. --go-grpc_out=. -I $protodir $protodir/entities/*.proto
protoc --go_out=. --go-grpc_out=. -I $protodir $protodir/services/*.proto
protoc --go_out=. --go-grpc_out=. -I $protodir $protodir/codehost/*proto

cp -R ./github.com/reviewpad/api/go/entities ./entities
cp -R ./github.com/reviewpad/api/go/services ./services
cp -R ./github.com/reviewpad/api/go/codehost ./codehost

rm -rf ./github.com
