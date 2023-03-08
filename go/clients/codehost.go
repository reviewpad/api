// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package clients

import (
	"github.com/reviewpad/api/go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCodeHostClient(endpoint string) (services.HostsClient, *grpc.ClientConn, error) {
	defaultOptions := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(419430400))
	transportCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())

	hostConnection, err := grpc.Dial(endpoint, transportCredentials, defaultOptions)
	if err != nil {
		return nil, nil, err
	}

	hostClient := services.NewHostsClient(hostConnection)

	return hostClient, hostConnection, nil
}