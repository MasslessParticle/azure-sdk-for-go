//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

const (
	moduleName    = "armhybridkubernetes"
	moduleVersion = "v0.2.0"
)

// AuthenticationMethod - The mode of client authentication.
type AuthenticationMethod string

const (
	AuthenticationMethodAAD   AuthenticationMethod = "AAD"
	AuthenticationMethodToken AuthenticationMethod = "Token"
)

// PossibleAuthenticationMethodValues returns the possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodAAD,
		AuthenticationMethodToken,
	}
}

// ToPtr returns a *AuthenticationMethod pointing to the current value.
func (c AuthenticationMethod) ToPtr() *AuthenticationMethod {
	return &c
}

// ConnectivityStatus - Represents the connectivity status of the connected cluster.
type ConnectivityStatus string

const (
	ConnectivityStatusConnected  ConnectivityStatus = "Connected"
	ConnectivityStatusConnecting ConnectivityStatus = "Connecting"
	ConnectivityStatusExpired    ConnectivityStatus = "Expired"
	ConnectivityStatusOffline    ConnectivityStatus = "Offline"
)

// PossibleConnectivityStatusValues returns the possible values for the ConnectivityStatus const type.
func PossibleConnectivityStatusValues() []ConnectivityStatus {
	return []ConnectivityStatus{
		ConnectivityStatusConnected,
		ConnectivityStatusConnecting,
		ConnectivityStatusExpired,
		ConnectivityStatusOffline,
	}
}

// ToPtr returns a *ConnectivityStatus pointing to the current value.
func (c ConnectivityStatus) ToPtr() *ConnectivityStatus {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// LastModifiedByType - The type of identity that last modified the resource.
type LastModifiedByType string

const (
	LastModifiedByTypeApplication     LastModifiedByType = "Application"
	LastModifiedByTypeKey             LastModifiedByType = "Key"
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	LastModifiedByTypeUser            LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns the possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{
		LastModifiedByTypeApplication,
		LastModifiedByTypeKey,
		LastModifiedByTypeManagedIdentity,
		LastModifiedByTypeUser,
	}
}

// ToPtr returns a *LastModifiedByType pointing to the current value.
func (c LastModifiedByType) ToPtr() *LastModifiedByType {
	return &c
}

// ProvisioningState - The current deployment state of connectedClusters.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// ResourceIdentityType - The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system
// created identity. The type 'None' means no identity is assigned to the connected cluster.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
	}
}

// ToPtr returns a *ResourceIdentityType pointing to the current value.
func (c ResourceIdentityType) ToPtr() *ResourceIdentityType {
	return &c
}
