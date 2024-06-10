// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apicenter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/services"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type ApiCenterServiceResource struct{}

func TestAccApiCenterService_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_center_service", "test")
	r := ApiCenterServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiCenterService_updateTags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_center_service", "test")
	r := ApiCenterServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateTags(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (ApiCenterServiceResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := services.ParseServiceID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ApiCenter.ServiceClient.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return pointer.To(resp.Model != nil && resp.Model.Id != nil), nil
}

func (r ApiCenterServiceResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_center_service" "test"{
	name                = "acctestapic-%d"
	location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

	tags = {
		acceptance = "Test"
	}
}
`, r.template(data), data.RandomInteger)
}

func (r ApiCenterServiceResource) updateTags(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_center_service" "test"{
	name                = "acctestapic-%d"
	location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

	tags = {
		acceptance = "Test"
    hello      = "World"
	}
}
`, r.template(data), data.RandomInteger)
}

func (r ApiCenterServiceResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
	features {}
}

resource "azurerm_resource_group" "test" {
	name     = "acctestRG-%d"
	location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}
