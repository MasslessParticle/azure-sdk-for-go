//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementpartner

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// Error - this is the management partner operations error
type Error struct {
	// this is the error response code
	Code *string `json:"code,omitempty"`

	// this is the ExtendedErrorInfo property
	Error *ExtendedErrorInfo `json:"error,omitempty"`

	// this is the extended error info message
	Message *string `json:"message,omitempty"`
}

// ExtendedErrorInfo - this is the extended error info
type ExtendedErrorInfo struct {
	// this is the error response code
	Code *string `json:"code,omitempty"`

	// this is the extended error info message
	Message *string `json:"message,omitempty"`
}

// OperationClientListOptions contains the optional parameters for the OperationClient.List method.
type OperationClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationDisplay - this is the management partner operation
type OperationDisplay struct {
	// the is management partner operation description
	Description *string `json:"description,omitempty"`

	// the is management partner operation
	Operation *string `json:"operation,omitempty"`

	// the is management partner provider
	Provider *string `json:"provider,omitempty"`

	// the is management partner resource
	Resource *string `json:"resource,omitempty"`
}

// OperationList - this is the management partner operations list
type OperationList struct {
	// Url to get the next page of items.
	NextLink *string `json:"nextLink,omitempty"`

	// this is the operation response list
	Value []*OperationResponse `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationResponse - this is the management partner operations response
type OperationResponse struct {
	// this is the operation display
	Display *OperationDisplay `json:"display,omitempty"`

	// this is the operation response name
	Name *string `json:"name,omitempty"`

	// the is operation response origin information
	Origin *string `json:"origin,omitempty"`
}

// PartnerClientCreateOptions contains the optional parameters for the PartnerClient.Create method.
type PartnerClientCreateOptions struct {
	// placeholder for future optional parameters
}

// PartnerClientDeleteOptions contains the optional parameters for the PartnerClient.Delete method.
type PartnerClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PartnerClientGetOptions contains the optional parameters for the PartnerClient.Get method.
type PartnerClientGetOptions struct {
	// placeholder for future optional parameters
}

// PartnerClientUpdateOptions contains the optional parameters for the PartnerClient.Update method.
type PartnerClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// PartnerProperties - this is the management partner properties
type PartnerProperties struct {
	// This is the DateTime when the partner was created.
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	// This is the object id.
	ObjectID *string `json:"objectId,omitempty"`

	// This is the partner id
	PartnerID *string `json:"partnerId,omitempty"`

	// This is the partner name
	PartnerName *string `json:"partnerName,omitempty"`

	// This is the partner state
	State *ManagementPartnerState `json:"state,omitempty"`

	// This is the tenant id.
	TenantID *string `json:"tenantId,omitempty"`

	// This is the DateTime when the partner was updated.
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`

	// This is the version.
	Version *int32 `json:"version,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PartnerProperties.
func (p PartnerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdTime", p.CreatedTime)
	populate(objectMap, "objectId", p.ObjectID)
	populate(objectMap, "partnerId", p.PartnerID)
	populate(objectMap, "partnerName", p.PartnerName)
	populate(objectMap, "state", p.State)
	populate(objectMap, "tenantId", p.TenantID)
	populateTimeRFC3339(objectMap, "updatedTime", p.UpdatedTime)
	populate(objectMap, "version", p.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PartnerProperties.
func (p *PartnerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdTime":
			err = unpopulateTimeRFC3339(val, &p.CreatedTime)
			delete(rawMsg, key)
		case "objectId":
			err = unpopulate(val, &p.ObjectID)
			delete(rawMsg, key)
		case "partnerId":
			err = unpopulate(val, &p.PartnerID)
			delete(rawMsg, key)
		case "partnerName":
			err = unpopulate(val, &p.PartnerName)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, &p.State)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, &p.TenantID)
			delete(rawMsg, key)
		case "updatedTime":
			err = unpopulateTimeRFC3339(val, &p.UpdatedTime)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &p.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PartnerResponse - this is the management partner operations response
type PartnerResponse struct {
	// Type of the partner
	Etag *int32 `json:"etag,omitempty"`

	// Properties of the partner
	Properties *PartnerProperties `json:"properties,omitempty"`

	// READ-ONLY; Identifier of the partner
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the partner
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of resource. "Microsoft.ManagementPartner/partners"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PartnersClientGetOptions contains the optional parameters for the PartnersClient.Get method.
type PartnersClientGetOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
