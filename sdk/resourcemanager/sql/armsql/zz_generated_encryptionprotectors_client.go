//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// EncryptionProtectorsClient contains the methods for the EncryptionProtectors group.
// Don't use this type directly, use NewEncryptionProtectorsClient() instead.
type EncryptionProtectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEncryptionProtectorsClient creates a new instance of EncryptionProtectorsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEncryptionProtectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EncryptionProtectorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &EncryptionProtectorsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// encryptionProtectorName - The name of the encryption protector to be updated.
// parameters - The requested encryption protector resource state.
// options - EncryptionProtectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the EncryptionProtectorsClient.BeginCreateOrUpdate
// method.
func (client *EncryptionProtectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsClientBeginCreateOrUpdateOptions) (EncryptionProtectorsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, encryptionProtectorName, parameters, options)
	if err != nil {
		return EncryptionProtectorsClientCreateOrUpdatePollerResponse{}, err
	}
	result := EncryptionProtectorsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EncryptionProtectorsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return EncryptionProtectorsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &EncryptionProtectorsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *EncryptionProtectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EncryptionProtectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets a server encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// encryptionProtectorName - The name of the encryption protector to be retrieved.
// options - EncryptionProtectorsClientGetOptions contains the optional parameters for the EncryptionProtectorsClient.Get
// method.
func (client *EncryptionProtectorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsClientGetOptions) (EncryptionProtectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return EncryptionProtectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EncryptionProtectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EncryptionProtectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EncryptionProtectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EncryptionProtectorsClient) getHandleResponse(resp *http.Response) (EncryptionProtectorsClientGetResponse, error) {
	result := EncryptionProtectorsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtector); err != nil {
		return EncryptionProtectorsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets a list of server encryption protectors
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - EncryptionProtectorsClientListByServerOptions contains the optional parameters for the EncryptionProtectorsClient.ListByServer
// method.
func (client *EncryptionProtectorsClient) ListByServer(resourceGroupName string, serverName string, options *EncryptionProtectorsClientListByServerOptions) *EncryptionProtectorsClientListByServerPager {
	return &EncryptionProtectorsClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp EncryptionProtectorsClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EncryptionProtectorListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *EncryptionProtectorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *EncryptionProtectorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *EncryptionProtectorsClient) listByServerHandleResponse(resp *http.Response) (EncryptionProtectorsClientListByServerResponse, error) {
	result := EncryptionProtectorsClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtectorListResult); err != nil {
		return EncryptionProtectorsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginRevalidate - Revalidates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// encryptionProtectorName - The name of the encryption protector to be updated.
// options - EncryptionProtectorsClientBeginRevalidateOptions contains the optional parameters for the EncryptionProtectorsClient.BeginRevalidate
// method.
func (client *EncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsClientBeginRevalidateOptions) (EncryptionProtectorsClientRevalidatePollerResponse, error) {
	resp, err := client.revalidate(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return EncryptionProtectorsClientRevalidatePollerResponse{}, err
	}
	result := EncryptionProtectorsClientRevalidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EncryptionProtectorsClient.Revalidate", "", resp, client.pl)
	if err != nil {
		return EncryptionProtectorsClientRevalidatePollerResponse{}, err
	}
	result.Poller = &EncryptionProtectorsClientRevalidatePoller{
		pt: pt,
	}
	return result, nil
}

// Revalidate - Revalidates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *EncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsClientBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *EncryptionProtectorsClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsClientBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
