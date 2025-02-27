//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import "net/http"

// LegacyPeeringsClientListResponse contains the response from method LegacyPeeringsClient.List.
type LegacyPeeringsClientListResponse struct {
	LegacyPeeringsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LegacyPeeringsClientListResult contains the result from method LegacyPeeringsClient.List.
type LegacyPeeringsClientListResult struct {
	ListResult
}

// LocationsClientListResponse contains the response from method LocationsClient.List.
type LocationsClientListResponse struct {
	LocationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationsClientListResult contains the result from method LocationsClient.List.
type LocationsClientListResult struct {
	LocationListResult
}

// ManagementClientCheckServiceProviderAvailabilityResponse contains the response from method ManagementClient.CheckServiceProviderAvailability.
type ManagementClientCheckServiceProviderAvailabilityResponse struct {
	ManagementClientCheckServiceProviderAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementClientCheckServiceProviderAvailabilityResult contains the result from method ManagementClient.CheckServiceProviderAvailability.
type ManagementClientCheckServiceProviderAvailabilityResult struct {
	Value *Enum0
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

// PeerAsnsClientCreateOrUpdateResponse contains the response from method PeerAsnsClient.CreateOrUpdate.
type PeerAsnsClientCreateOrUpdateResponse struct {
	PeerAsnsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeerAsnsClientCreateOrUpdateResult contains the result from method PeerAsnsClient.CreateOrUpdate.
type PeerAsnsClientCreateOrUpdateResult struct {
	PeerAsn
}

// PeerAsnsClientDeleteResponse contains the response from method PeerAsnsClient.Delete.
type PeerAsnsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeerAsnsClientGetResponse contains the response from method PeerAsnsClient.Get.
type PeerAsnsClientGetResponse struct {
	PeerAsnsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeerAsnsClientGetResult contains the result from method PeerAsnsClient.Get.
type PeerAsnsClientGetResult struct {
	PeerAsn
}

// PeerAsnsClientListBySubscriptionResponse contains the response from method PeerAsnsClient.ListBySubscription.
type PeerAsnsClientListBySubscriptionResponse struct {
	PeerAsnsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeerAsnsClientListBySubscriptionResult contains the result from method PeerAsnsClient.ListBySubscription.
type PeerAsnsClientListBySubscriptionResult struct {
	PeerAsnListResult
}

// PeeringsClientCreateOrUpdateResponse contains the response from method PeeringsClient.CreateOrUpdate.
type PeeringsClientCreateOrUpdateResponse struct {
	PeeringsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientCreateOrUpdateResult contains the result from method PeeringsClient.CreateOrUpdate.
type PeeringsClientCreateOrUpdateResult struct {
	Peering
}

// PeeringsClientDeleteResponse contains the response from method PeeringsClient.Delete.
type PeeringsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientGetResponse contains the response from method PeeringsClient.Get.
type PeeringsClientGetResponse struct {
	PeeringsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientGetResult contains the result from method PeeringsClient.Get.
type PeeringsClientGetResult struct {
	Peering
}

// PeeringsClientListByResourceGroupResponse contains the response from method PeeringsClient.ListByResourceGroup.
type PeeringsClientListByResourceGroupResponse struct {
	PeeringsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientListByResourceGroupResult contains the result from method PeeringsClient.ListByResourceGroup.
type PeeringsClientListByResourceGroupResult struct {
	ListResult
}

// PeeringsClientListBySubscriptionResponse contains the response from method PeeringsClient.ListBySubscription.
type PeeringsClientListBySubscriptionResponse struct {
	PeeringsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientListBySubscriptionResult contains the result from method PeeringsClient.ListBySubscription.
type PeeringsClientListBySubscriptionResult struct {
	ListResult
}

// PeeringsClientUpdateResponse contains the response from method PeeringsClient.Update.
type PeeringsClientUpdateResponse struct {
	PeeringsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PeeringsClientUpdateResult contains the result from method PeeringsClient.Update.
type PeeringsClientUpdateResult struct {
	Peering
}

// PrefixesClientListByPeeringServiceResponse contains the response from method PrefixesClient.ListByPeeringService.
type PrefixesClientListByPeeringServiceResponse struct {
	PrefixesClientListByPeeringServiceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrefixesClientListByPeeringServiceResult contains the result from method PrefixesClient.ListByPeeringService.
type PrefixesClientListByPeeringServiceResult struct {
	ServicePrefixListResult
}

// ServiceLocationsClientListResponse contains the response from method ServiceLocationsClient.List.
type ServiceLocationsClientListResponse struct {
	ServiceLocationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServiceLocationsClientListResult contains the result from method ServiceLocationsClient.List.
type ServiceLocationsClientListResult struct {
	ServiceLocationListResult
}

// ServicePrefixesClientCreateOrUpdateResponse contains the response from method ServicePrefixesClient.CreateOrUpdate.
type ServicePrefixesClientCreateOrUpdateResponse struct {
	ServicePrefixesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicePrefixesClientCreateOrUpdateResult contains the result from method ServicePrefixesClient.CreateOrUpdate.
type ServicePrefixesClientCreateOrUpdateResult struct {
	ServicePrefix
}

// ServicePrefixesClientDeleteResponse contains the response from method ServicePrefixesClient.Delete.
type ServicePrefixesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicePrefixesClientGetResponse contains the response from method ServicePrefixesClient.Get.
type ServicePrefixesClientGetResponse struct {
	ServicePrefixesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicePrefixesClientGetResult contains the result from method ServicePrefixesClient.Get.
type ServicePrefixesClientGetResult struct {
	ServicePrefix
}

// ServiceProvidersClientListResponse contains the response from method ServiceProvidersClient.List.
type ServiceProvidersClientListResponse struct {
	ServiceProvidersClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServiceProvidersClientListResult contains the result from method ServiceProvidersClient.List.
type ServiceProvidersClientListResult struct {
	ServiceProviderListResult
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServicesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientCreateOrUpdateResult contains the result from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResult struct {
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServicesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientGetResult contains the result from method ServicesClient.Get.
type ServicesClientGetResult struct {
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResponse struct {
	ServicesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientListByResourceGroupResult contains the result from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResult struct {
	ServiceListResult
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.ListBySubscription.
type ServicesClientListBySubscriptionResponse struct {
	ServicesClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientListBySubscriptionResult contains the result from method ServicesClient.ListBySubscription.
type ServicesClientListBySubscriptionResult struct {
	ServiceListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServicesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientUpdateResult contains the result from method ServicesClient.Update.
type ServicesClientUpdateResult struct {
	Service
}
