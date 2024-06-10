// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/apis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/workspaces"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ApiClient       *apis.ApisClient
	ServiceClient   *services.ServicesClient
	WorkspaceClient *workspaces.WorkspacesClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	apiClient, err := apis.NewApisClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Api Clinet: %+v", err)
	}
	o.Configure(apiClient.Client, o.Authorizers.ResourceManager)

	serviceClient, err := services.NewServicesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Service Client: %+v", err)
	}
	o.Configure(serviceClient.Client, o.Authorizers.ResourceManager)

	workspaceClient, err := workspaces.NewWorkspacesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Workspace Client: %+v", err)
	}
	o.Configure(workspaceClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		ApiClient:       apiClient,
		ServiceClient:   serviceClient,
		WorkspaceClient: workspaceClient,
	}, nil
}
