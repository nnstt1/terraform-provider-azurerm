// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apicenter

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
)

type Registration struct{}

var _ sdk.TypedServiceRegistration = Registration{}

func (r Registration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{}
}

func (r Registration) Resources() []sdk.Resource {
	return []sdk.Resource{
		ApiCenterServiceResource{},
	}
}

// Name is the name of this Service
func (r Registration) Name() string {
	return "API Center"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"API Center",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
// func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
// 	return map[string]*pluginsdk.Resource{
// 		"azurerm_api_management":                                 dataSourceApiManagementService(),
// 	}
// }

// SupportedResources returns the supported Resources supported by this Service
// func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
// 	return map[string]*pluginsdk.Resource{
// 		"azurerm_api_management":                                 resourceApiManagementService(),
// 	}
// }

// func (r Registration) AssociatedGitHubLabel() string {
// 	return "service/api-center"
// }
