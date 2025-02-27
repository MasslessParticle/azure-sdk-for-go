//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DashboardsClient contains the methods for the Dashboards group.
// Don't use this type directly, use NewDashboardsClient() instead.
type DashboardsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDashboardsClient creates a new instance of DashboardsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDashboardsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DashboardsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DashboardsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a Dashboard.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// dashboardName - The name of the dashboard.
// dashboard - The parameters required to create or update a dashboard.
// options - DashboardsClientCreateOrUpdateOptions contains the optional parameters for the DashboardsClient.CreateOrUpdate
// method.
func (client *DashboardsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dashboardName string, dashboard Dashboard, options *DashboardsClientCreateOrUpdateOptions) (DashboardsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dashboardName, dashboard, options)
	if err != nil {
		return DashboardsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DashboardsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DashboardsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DashboardsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dashboardName string, dashboard Dashboard, options *DashboardsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dashboardName == "" {
		return nil, errors.New("parameter dashboardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dashboardName}", url.PathEscape(dashboardName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dashboard)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DashboardsClient) createOrUpdateHandleResponse(resp *http.Response) (DashboardsClientCreateOrUpdateResponse, error) {
	result := DashboardsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Dashboard); err != nil {
		return DashboardsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the Dashboard.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// dashboardName - The name of the dashboard.
// options - DashboardsClientDeleteOptions contains the optional parameters for the DashboardsClient.Delete method.
func (client *DashboardsClient) Delete(ctx context.Context, resourceGroupName string, dashboardName string, options *DashboardsClientDeleteOptions) (DashboardsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dashboardName, options)
	if err != nil {
		return DashboardsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DashboardsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DashboardsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DashboardsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DashboardsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dashboardName string, options *DashboardsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dashboardName == "" {
		return nil, errors.New("parameter dashboardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dashboardName}", url.PathEscape(dashboardName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the Dashboard.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// dashboardName - The name of the dashboard.
// options - DashboardsClientGetOptions contains the optional parameters for the DashboardsClient.Get method.
func (client *DashboardsClient) Get(ctx context.Context, resourceGroupName string, dashboardName string, options *DashboardsClientGetOptions) (DashboardsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dashboardName, options)
	if err != nil {
		return DashboardsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DashboardsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return DashboardsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DashboardsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dashboardName string, options *DashboardsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dashboardName == "" {
		return nil, errors.New("parameter dashboardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dashboardName}", url.PathEscape(dashboardName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DashboardsClient) getHandleResponse(resp *http.Response) (DashboardsClientGetResponse, error) {
	result := DashboardsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Dashboard); err != nil {
		return DashboardsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all the Dashboards within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - DashboardsClientListByResourceGroupOptions contains the optional parameters for the DashboardsClient.ListByResourceGroup
// method.
func (client *DashboardsClient) ListByResourceGroup(resourceGroupName string, options *DashboardsClientListByResourceGroupOptions) *DashboardsClientListByResourceGroupPager {
	return &DashboardsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DashboardsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DashboardListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DashboardsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DashboardsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DashboardsClient) listByResourceGroupHandleResponse(resp *http.Response) (DashboardsClientListByResourceGroupResponse, error) {
	result := DashboardsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DashboardListResult); err != nil {
		return DashboardsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets all the dashboards within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DashboardsClientListBySubscriptionOptions contains the optional parameters for the DashboardsClient.ListBySubscription
// method.
func (client *DashboardsClient) ListBySubscription(options *DashboardsClientListBySubscriptionOptions) *DashboardsClientListBySubscriptionPager {
	return &DashboardsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DashboardsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DashboardListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DashboardsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DashboardsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Portal/dashboards"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DashboardsClient) listBySubscriptionHandleResponse(resp *http.Response) (DashboardsClientListBySubscriptionResponse, error) {
	result := DashboardsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DashboardListResult); err != nil {
		return DashboardsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Dashboard.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// dashboardName - The name of the dashboard.
// dashboard - The updatable fields of a Dashboard.
// options - DashboardsClientUpdateOptions contains the optional parameters for the DashboardsClient.Update method.
func (client *DashboardsClient) Update(ctx context.Context, resourceGroupName string, dashboardName string, dashboard PatchableDashboard, options *DashboardsClientUpdateOptions) (DashboardsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dashboardName, dashboard, options)
	if err != nil {
		return DashboardsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DashboardsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return DashboardsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DashboardsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dashboardName string, dashboard PatchableDashboard, options *DashboardsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dashboardName == "" {
		return nil, errors.New("parameter dashboardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dashboardName}", url.PathEscape(dashboardName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dashboard)
}

// updateHandleResponse handles the Update response.
func (client *DashboardsClient) updateHandleResponse(resp *http.Response) (DashboardsClientUpdateResponse, error) {
	result := DashboardsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Dashboard); err != nil {
		return DashboardsClientUpdateResponse{}, err
	}
	return result, nil
}
