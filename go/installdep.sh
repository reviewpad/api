#!/bin/bash -eu
# Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
# Use of this source code is governed by a license that can be
# found in the LICENSE file.

# Install protoc
PROTOC_ZIP=protoc-22.2-osx-x86_64.zip

curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v22.2/$PROTOC_ZIP && \
            unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
            unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
            rm -f $PROTOC_ZIP

go install google.golang.org/protobuf/cmd/protoc-gen-go
