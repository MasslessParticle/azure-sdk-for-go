//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceupdate

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AccountsClientCreatePollerResponse contains the response from method AccountsClient.Create.
type AccountsClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientCreateResponse, error) {
	respType := AccountsClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientCreatePollerResponse from the provided client and resume token.
func (l *AccountsClientCreatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	AccountsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientCreateResult contains the result from method AccountsClient.Create.
type AccountsClientCreateResult struct {
	Account
}

// AccountsClientDeletePollerResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientDeleteResponse, error) {
	respType := AccountsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientDeletePollerResponse from the provided client and resume token.
func (l *AccountsClientDeletePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	AccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientGetResult contains the result from method AccountsClient.Get.
type AccountsClientGetResult struct {
	Account
}

// AccountsClientHeadResponse contains the response from method AccountsClient.Head.
type AccountsClientHeadResponse struct {
	AccountsClientHeadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientHeadResult contains the result from method AccountsClient.Head.
type AccountsClientHeadResult struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResponse struct {
	AccountsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListByResourceGroupResult contains the result from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResult struct {
	AccountList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.ListBySubscription.
type AccountsClientListBySubscriptionResponse struct {
	AccountsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListBySubscriptionResult contains the result from method AccountsClient.ListBySubscription.
type AccountsClientListBySubscriptionResult struct {
	AccountList
}

// AccountsClientUpdatePollerResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientUpdateResponse, error) {
	respType := AccountsClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientUpdatePollerResponse from the provided client and resume token.
func (l *AccountsClientUpdatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	AccountsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientUpdateResult contains the result from method AccountsClient.Update.
type AccountsClientUpdateResult struct {
	Account
}

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	ClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCheckNameAvailabilityResult contains the result from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResponse
}

// InstancesClientCreatePollerResponse contains the response from method InstancesClient.Create.
type InstancesClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *InstancesClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l InstancesClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (InstancesClientCreateResponse, error) {
	respType := InstancesClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Instance)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a InstancesClientCreatePollerResponse from the provided client and resume token.
func (l *InstancesClientCreatePollerResponse) Resume(ctx context.Context, client *InstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("InstancesClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &InstancesClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// InstancesClientCreateResponse contains the response from method InstancesClient.Create.
type InstancesClientCreateResponse struct {
	InstancesClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientCreateResult contains the result from method InstancesClient.Create.
type InstancesClientCreateResult struct {
	Instance
}

// InstancesClientDeletePollerResponse contains the response from method InstancesClient.Delete.
type InstancesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *InstancesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l InstancesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (InstancesClientDeleteResponse, error) {
	respType := InstancesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a InstancesClientDeletePollerResponse from the provided client and resume token.
func (l *InstancesClientDeletePollerResponse) Resume(ctx context.Context, client *InstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("InstancesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &InstancesClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// InstancesClientDeleteResponse contains the response from method InstancesClient.Delete.
type InstancesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientGetResponse contains the response from method InstancesClient.Get.
type InstancesClientGetResponse struct {
	InstancesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientGetResult contains the result from method InstancesClient.Get.
type InstancesClientGetResult struct {
	Instance
}

// InstancesClientHeadResponse contains the response from method InstancesClient.Head.
type InstancesClientHeadResponse struct {
	InstancesClientHeadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientHeadResult contains the result from method InstancesClient.Head.
type InstancesClientHeadResult struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// InstancesClientListByAccountResponse contains the response from method InstancesClient.ListByAccount.
type InstancesClientListByAccountResponse struct {
	InstancesClientListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientListByAccountResult contains the result from method InstancesClient.ListByAccount.
type InstancesClientListByAccountResult struct {
	InstanceList
}

// InstancesClientUpdateResponse contains the response from method InstancesClient.Update.
type InstancesClientUpdateResponse struct {
	InstancesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InstancesClientUpdateResult contains the result from method InstancesClient.Update.
type InstancesClientUpdateResult struct {
	Instance
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionProxiesClientCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnectionProxiesClient.CreateOrUpdate.
type PrivateEndpointConnectionProxiesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionProxiesClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionProxiesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnectionProxy)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionProxiesClientCreateOrUpdatePollerResponse from the provided client and resume
// token.
func (l *PrivateEndpointConnectionProxiesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionProxiesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionProxiesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionProxiesClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionProxiesClient.CreateOrUpdate.
type PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionProxiesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionProxiesClientCreateOrUpdateResult contains the result from method PrivateEndpointConnectionProxiesClient.CreateOrUpdate.
type PrivateEndpointConnectionProxiesClientCreateOrUpdateResult struct {
	PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxiesClientDeletePollerResponse contains the response from method PrivateEndpointConnectionProxiesClient.Delete.
type PrivateEndpointConnectionProxiesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionProxiesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionProxiesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionProxiesClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionProxiesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionProxiesClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionProxiesClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionProxiesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionProxiesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionProxiesClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionProxiesClientDeleteResponse contains the response from method PrivateEndpointConnectionProxiesClient.Delete.
type PrivateEndpointConnectionProxiesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionProxiesClientGetResponse contains the response from method PrivateEndpointConnectionProxiesClient.Get.
type PrivateEndpointConnectionProxiesClientGetResponse struct {
	PrivateEndpointConnectionProxiesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionProxiesClientGetResult contains the result from method PrivateEndpointConnectionProxiesClient.Get.
type PrivateEndpointConnectionProxiesClientGetResult struct {
	PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxiesClientListByAccountResponse contains the response from method PrivateEndpointConnectionProxiesClient.ListByAccount.
type PrivateEndpointConnectionProxiesClientListByAccountResponse struct {
	PrivateEndpointConnectionProxiesClientListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionProxiesClientListByAccountResult contains the result from method PrivateEndpointConnectionProxiesClient.ListByAccount.
type PrivateEndpointConnectionProxiesClientListByAccountResult struct {
	PrivateEndpointConnectionProxyListResult
}

// PrivateEndpointConnectionProxiesClientValidateResponse contains the response from method PrivateEndpointConnectionProxiesClient.Validate.
type PrivateEndpointConnectionProxiesClientValidateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientCreateOrUpdateResult contains the result from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeletePollerResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnectionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResult contains the result from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByAccountResponse contains the response from method PrivateEndpointConnectionsClient.ListByAccount.
type PrivateEndpointConnectionsClientListByAccountResponse struct {
	PrivateEndpointConnectionsClientListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientListByAccountResult contains the result from method PrivateEndpointConnectionsClient.ListByAccount.
type PrivateEndpointConnectionsClientListByAccountResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourcesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientGetResult contains the result from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResult struct {
	GroupInformation
}

// PrivateLinkResourcesClientListByAccountResponse contains the response from method PrivateLinkResourcesClient.ListByAccount.
type PrivateLinkResourcesClientListByAccountResponse struct {
	PrivateLinkResourcesClientListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientListByAccountResult contains the result from method PrivateLinkResourcesClient.ListByAccount.
type PrivateLinkResourcesClientListByAccountResult struct {
	PrivateLinkResourceListResult
}
