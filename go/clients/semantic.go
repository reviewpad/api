// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package clients

import (
	"crypto/tls"

	"github.com/reviewpad/api/go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewSemanticClient(endpoint string) (services.SemanticClient, *grpc.ClientConn, error) {
	defaultOptions := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(419430400))
	transportCredentials := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))

	semanticConnection, err := grpc.Dial(endpoint, transportCredentials, defaultOptions)
	if err != nil {
		return nil, nil, err
	}

	semanticClient := services.NewSemanticClient(semanticConnection)

	return semanticClient, semanticConnection, nil
}
