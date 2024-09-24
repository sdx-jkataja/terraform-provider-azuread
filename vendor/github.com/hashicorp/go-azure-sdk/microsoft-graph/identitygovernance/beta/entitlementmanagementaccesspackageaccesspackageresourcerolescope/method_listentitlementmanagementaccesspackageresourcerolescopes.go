package entitlementmanagementaccesspackageaccesspackageresourcerolescope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageResourceRoleScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions() ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions {
	return ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopes - Get accessPackageResourceRoleScopes from
// identityGovernance
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopes(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions) (result ListEntitlementManagementAccessPackageResourceRoleScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResourceRoleScopes", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.AccessPackageResourceRoleScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageResourceRoleScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions) (ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageResourceRoleScopesOperationOptions, predicate AccessPackageResourceRoleScopeOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceRoleScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopes(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
