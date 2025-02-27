//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineList.json
func ExampleRulesEnginesClient_ListByFrontDoor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	pager := client.ListByFrontDoor("<resource-group-name>",
		"<front-door-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineGet.json
func ExampleRulesEnginesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RulesEnginesClientGetResult)
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineCreate.json
func ExampleRulesEnginesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		armfrontdoor.RulesEngine{
			Properties: &armfrontdoor.RulesEngineProperties{
				Rules: []*armfrontdoor.RulesEngineRule{
					{
						Name: to.StringPtr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
								ODataType:         to.StringPtr("<odata-type>"),
								CustomFragment:    to.StringPtr("<custom-fragment>"),
								CustomHost:        to.StringPtr("<custom-host>"),
								CustomPath:        to.StringPtr("<custom-path>"),
								CustomQueryString: to.StringPtr("<custom-query-string>"),
								RedirectProtocol:  armfrontdoor.FrontDoorRedirectProtocol("HttpsOnly").ToPtr(),
								RedirectType:      armfrontdoor.FrontDoorRedirectType("Moved").ToPtr(),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.StringPtr("CH")},
								RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariable("RemoteAddr").ToPtr(),
								RulesEngineOperator:      armfrontdoor.RulesEngineOperator("GeoMatch").ToPtr(),
							}},
						MatchProcessingBehavior: armfrontdoor.MatchProcessingBehavior("Stop").ToPtr(),
						Priority:                to.Int32Ptr(1),
					},
					{
						Name: to.StringPtr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							ResponseHeaderActions: []*armfrontdoor.HeaderAction{
								{
									HeaderActionType: armfrontdoor.HeaderActionType("Overwrite").ToPtr(),
									HeaderName:       to.StringPtr("<header-name>"),
									Value:            to.StringPtr("<value>"),
								}},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.StringPtr("jpg")},
								RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariable("RequestFilenameExtension").ToPtr(),
								RulesEngineOperator:      armfrontdoor.RulesEngineOperator("Equal").ToPtr(),
								Transforms: []*armfrontdoor.Transform{
									armfrontdoor.Transform("Lowercase").ToPtr()},
							}},
						Priority: to.Int32Ptr(2),
					},
					{
						Name: to.StringPtr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.StringPtr("<odata-type>"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.StringPtr("<id>"),
								},
								CacheConfiguration: &armfrontdoor.CacheConfiguration{
									CacheDuration:                to.StringPtr("<cache-duration>"),
									DynamicCompression:           armfrontdoor.DynamicCompressionEnabled("Disabled").ToPtr(),
									QueryParameterStripDirective: armfrontdoor.FrontDoorQuery("StripOnly").ToPtr(),
									QueryParameters:              to.StringPtr("<query-parameters>"),
								},
								ForwardingProtocol: armfrontdoor.FrontDoorForwardingProtocol("HttpsOnly").ToPtr(),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								NegateCondition: to.BoolPtr(false),
								RulesEngineMatchValue: []*string{
									to.StringPtr("allowoverride")},
								RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariable("RequestHeader").ToPtr(),
								RulesEngineOperator:      armfrontdoor.RulesEngineOperator("Equal").ToPtr(),
								Selector:                 to.StringPtr("<selector>"),
								Transforms: []*armfrontdoor.Transform{
									armfrontdoor.Transform("Lowercase").ToPtr()},
							}},
						Priority: to.Int32Ptr(3),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RulesEnginesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineDelete.json
func ExampleRulesEnginesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
