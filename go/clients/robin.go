// Copyright (C) 2023 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package clients

import (
	"crypto/tls"

	"github.com/reviewpad/api/go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewRobinClient(endpoint string) (services.RobinClient, *grpc.ClientConn, error) {
	defaultOptions := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(419430400))
	transportCredentials := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))

	robinConnection, err := grpc.Dial(endpoint, transportCredentials, defaultOptions)
	if err != nil {
		return nil, nil, err
	}

	robinClient := services.NewRobinClient(robinConnection)

	return robinClient, robinConnection, nil
}
