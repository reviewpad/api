// Copyright (C) 2023 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package clients

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v49/github"
)

// GitHubAppClient is a client for the GitHub App API.
type GitHubAppClient struct {
	client *github.AppsService
}

// NewGitHubAppClient creates a new GitHubAppClient.
// appId is the GitHub App ID.
// appPrivateKey is the private key of the GitHub App.
func NewGitHubAppClient(appId int64, appPrivateKey []byte) (*GitHubAppClient, error) {
	transport, err := ghinstallation.NewAppsTransport(http.DefaultTransport, appId, appPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &GitHubAppClient{
		client: github.NewClient(&http.Client{Transport: transport}).Apps,
	}, nil
}

// CreateInstallationToken creates a new installation token for the given installation ID.
func (c *GitHubAppClient) CreateInstallationToken(ctx context.Context, installationID int64) (string, error) {
	token, _, err := c.client.CreateInstallationToken(ctx, installationID, &github.InstallationTokenOptions{})
	if err != nil {
		return "", fmt.Errorf("error creating installation token: %v", err)
	}

	return token.GetToken(), nil
}
