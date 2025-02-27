# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ApplicationDefinitionsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *ApplicationDefinitionsBeginDeleteOptions)` to `(context.Context, string, string, *ApplicationDefinitionsClientBeginDeleteOptions)`
- Function `*ApplicationDefinitionsClient.BeginDelete` return value(s) have been changed from `(ApplicationDefinitionsDeletePollerResponse, error)` to `(ApplicationDefinitionsClientDeletePollerResponse, error)`
- Function `*ApplicationDefinitionsClient.Update` parameter(s) have been changed from `(context.Context, string, string, ApplicationDefinitionPatchable, *ApplicationDefinitionsUpdateOptions)` to `(context.Context, string, string, ApplicationDefinitionPatchable, *ApplicationDefinitionsClientUpdateOptions)`
- Function `*ApplicationDefinitionsClient.Update` return value(s) have been changed from `(ApplicationDefinitionsUpdateResponse, error)` to `(ApplicationDefinitionsClientUpdateResponse, error)`
- Function `*JitRequestsClient.ListByResourceGroup` parameter(s) have been changed from `(context.Context, string, *JitRequestsListByResourceGroupOptions)` to `(context.Context, string, *JitRequestsClientListByResourceGroupOptions)`
- Function `*JitRequestsClient.ListByResourceGroup` return value(s) have been changed from `(JitRequestsListByResourceGroupResponse, error)` to `(JitRequestsClientListByResourceGroupResponse, error)`
- Function `*ApplicationDefinitionsClient.ListBySubscription` parameter(s) have been changed from `(*ApplicationDefinitionsListBySubscriptionOptions)` to `(*ApplicationDefinitionsClientListBySubscriptionOptions)`
- Function `*ApplicationDefinitionsClient.ListBySubscription` return value(s) have been changed from `(*ApplicationDefinitionsListBySubscriptionPager)` to `(*ApplicationDefinitionsClientListBySubscriptionPager)`
- Function `*JitRequestsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, *JitRequestsDeleteOptions)` to `(context.Context, string, string, *JitRequestsClientDeleteOptions)`
- Function `*JitRequestsClient.Delete` return value(s) have been changed from `(JitRequestsDeleteResponse, error)` to `(JitRequestsClientDeleteResponse, error)`
- Function `*ApplicationDefinitionsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *ApplicationDefinitionsListByResourceGroupOptions)` to `(string, *ApplicationDefinitionsClientListByResourceGroupOptions)`
- Function `*ApplicationDefinitionsClient.ListByResourceGroup` return value(s) have been changed from `(*ApplicationDefinitionsListByResourceGroupPager)` to `(*ApplicationDefinitionsClientListByResourceGroupPager)`
- Function `*ApplicationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsGetOptions)` to `(context.Context, string, string, *ApplicationsClientGetOptions)`
- Function `*ApplicationsClient.Get` return value(s) have been changed from `(ApplicationsGetResponse, error)` to `(ApplicationsClientGetResponse, error)`
- Function `*ApplicationsClient.ListAllowedUpgradePlans` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsListAllowedUpgradePlansOptions)` to `(context.Context, string, string, *ApplicationsClientListAllowedUpgradePlansOptions)`
- Function `*ApplicationsClient.ListAllowedUpgradePlans` return value(s) have been changed from `(ApplicationsListAllowedUpgradePlansResponse, error)` to `(ApplicationsClientListAllowedUpgradePlansResponse, error)`
- Function `*ApplicationDefinitionsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *ApplicationDefinitionsGetOptions)` to `(context.Context, string, string, *ApplicationDefinitionsClientGetOptions)`
- Function `*ApplicationDefinitionsClient.Get` return value(s) have been changed from `(ApplicationDefinitionsGetResponse, error)` to `(ApplicationDefinitionsClientGetResponse, error)`
- Function `*JitRequestsClient.ListBySubscription` parameter(s) have been changed from `(context.Context, *JitRequestsListBySubscriptionOptions)` to `(context.Context, *JitRequestsClientListBySubscriptionOptions)`
- Function `*JitRequestsClient.ListBySubscription` return value(s) have been changed from `(JitRequestsListBySubscriptionResponse, error)` to `(JitRequestsClientListBySubscriptionResponse, error)`
- Function `*ApplicationsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *ApplicationsListByResourceGroupOptions)` to `(string, *ApplicationsClientListByResourceGroupOptions)`
- Function `*ApplicationsClient.ListByResourceGroup` return value(s) have been changed from `(*ApplicationsListByResourceGroupPager)` to `(*ApplicationsClientListByResourceGroupPager)`
- Function `*JitRequestsClient.Update` parameter(s) have been changed from `(context.Context, string, string, JitRequestPatchable, *JitRequestsUpdateOptions)` to `(context.Context, string, string, JitRequestPatchable, *JitRequestsClientUpdateOptions)`
- Function `*JitRequestsClient.Update` return value(s) have been changed from `(JitRequestsUpdateResponse, error)` to `(JitRequestsClientUpdateResponse, error)`
- Function `*ApplicationsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, Application, *ApplicationsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, Application, *ApplicationsClientBeginCreateOrUpdateOptions)`
- Function `*ApplicationsClient.BeginCreateOrUpdate` return value(s) have been changed from `(ApplicationsCreateOrUpdatePollerResponse, error)` to `(ApplicationsClientCreateOrUpdatePollerResponse, error)`
- Function `*ApplicationsClient.BeginRefreshPermissions` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsBeginRefreshPermissionsOptions)` to `(context.Context, string, string, *ApplicationsClientBeginRefreshPermissionsOptions)`
- Function `*ApplicationsClient.BeginRefreshPermissions` return value(s) have been changed from `(ApplicationsRefreshPermissionsPollerResponse, error)` to `(ApplicationsClientRefreshPermissionsPollerResponse, error)`
- Function `*ApplicationsClient.ListBySubscription` parameter(s) have been changed from `(*ApplicationsListBySubscriptionOptions)` to `(*ApplicationsClientListBySubscriptionOptions)`
- Function `*ApplicationsClient.ListBySubscription` return value(s) have been changed from `(*ApplicationsListBySubscriptionPager)` to `(*ApplicationsClientListBySubscriptionPager)`
- Function `*JitRequestsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, JitRequestDefinition, *JitRequestsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, JitRequestDefinition, *JitRequestsClientBeginCreateOrUpdateOptions)`
- Function `*JitRequestsClient.BeginCreateOrUpdate` return value(s) have been changed from `(JitRequestsCreateOrUpdatePollerResponse, error)` to `(JitRequestsClientCreateOrUpdatePollerResponse, error)`
- Function `*ApplicationDefinitionsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, ApplicationDefinition, *ApplicationDefinitionsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, ApplicationDefinition, *ApplicationDefinitionsClientBeginCreateOrUpdateOptions)`
- Function `*ApplicationDefinitionsClient.BeginCreateOrUpdate` return value(s) have been changed from `(ApplicationDefinitionsCreateOrUpdatePollerResponse, error)` to `(ApplicationDefinitionsClientCreateOrUpdatePollerResponse, error)`
- Function `*ApplicationsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsUpdateOptions)` to `(context.Context, string, string, *ApplicationsClientUpdateOptions)`
- Function `*ApplicationsClient.Update` return value(s) have been changed from `(ApplicationsUpdateResponse, error)` to `(ApplicationsClientUpdateResponse, error)`
- Function `*ApplicationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsBeginDeleteOptions)` to `(context.Context, string, string, *ApplicationsClientBeginDeleteOptions)`
- Function `*ApplicationsClient.BeginDelete` return value(s) have been changed from `(ApplicationsDeletePollerResponse, error)` to `(ApplicationsClientDeletePollerResponse, error)`
- Function `*JitRequestsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *JitRequestsGetOptions)` to `(context.Context, string, string, *JitRequestsClientGetOptions)`
- Function `*JitRequestsClient.Get` return value(s) have been changed from `(JitRequestsGetResponse, error)` to `(JitRequestsClientGetResponse, error)`
- Type of `Application.Identity` has been changed from `*ManagedServiceIdentity` to `*Identity`
- Type of `ApplicationPatchable.Identity` has been changed from `*ManagedServiceIdentity` to `*Identity`
- Const `ManagedServiceIdentityTypeUserAssigned` has been removed
- Const `ManagedServiceIdentityTypeSystemAssignedUserAssigned` has been removed
- Const `ManagedServiceIdentityTypeNone` has been removed
- Const `ManagedServiceIdentityTypeSystemAssigned` has been removed
- Function `*ApplicationDefinitionsListByResourceGroupPager.PageResponse` has been removed
- Function `*ApplicationsListByResourceGroupPager.PageResponse` has been removed
- Function `ApplicationsRefreshPermissionsPollerResponse.PollUntilDone` has been removed
- Function `*ApplicationDefinitionsListByResourceGroupPager.Err` has been removed
- Function `*JitRequestsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*ApplicationsListByResourceGroupPager.Err` has been removed
- Function `*ApplicationDefinitionsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*ApplicationDefinitionsListBySubscriptionPager.NextPage` has been removed
- Function `*JitRequestsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `JitRequestsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ApplicationsListBySubscriptionPager.PageResponse` has been removed
- Function `*ApplicationDefinitionsListByResourceGroupPager.NextPage` has been removed
- Function `*ApplicationDefinitionsDeletePoller.Done` has been removed
- Function `*ApplicationDefinitionsDeletePoller.ResumeToken` has been removed
- Function `*ApplicationDefinitionsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*ApplicationsDeletePoller.ResumeToken` has been removed
- Function `*JitRequestsCreateOrUpdatePoller.Done` has been removed
- Function `*ApplicationsRefreshPermissionsPoller.Poll` has been removed
- Function `*ApplicationDefinitionsCreateOrUpdatePoller.Poll` has been removed
- Function `*ApplicationsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*ApplicationDefinitionsDeletePoller.FinalResponse` has been removed
- Function `*ApplicationsRefreshPermissionsPollerResponse.Resume` has been removed
- Function `ApplicationDefinitionsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ApplicationsDeletePollerResponse.Resume` has been removed
- Function `*ApplicationDefinitionsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*ApplicationsCreateOrUpdatePoller.Poll` has been removed
- Function `*ApplicationDefinitionsCreateOrUpdatePoller.Done` has been removed
- Function `ApplicationsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ApplicationsListBySubscriptionPager.NextPage` has been removed
- Function `*ApplicationsDeletePoller.Done` has been removed
- Function `*ApplicationsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*ApplicationDefinitionsListBySubscriptionPager.Err` has been removed
- Function `*ApplicationsCreateOrUpdatePoller.Done` has been removed
- Function `*ApplicationsRefreshPermissionsPoller.Done` has been removed
- Function `ApplicationDefinitionsDeletePollerResponse.PollUntilDone` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*JitRequestsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `ApplicationsDeletePollerResponse.PollUntilDone` has been removed
- Function `*ApplicationsDeletePoller.Poll` has been removed
- Function `ManagedServiceIdentity.MarshalJSON` has been removed
- Function `*ApplicationDefinitionsDeletePoller.Poll` has been removed
- Function `ManagedServiceIdentityType.ToPtr` has been removed
- Function `*ApplicationsDeletePoller.FinalResponse` has been removed
- Function `*ApplicationDefinitionsDeletePollerResponse.Resume` has been removed
- Function `*ApplicationsRefreshPermissionsPoller.ResumeToken` has been removed
- Function `*ApplicationsRefreshPermissionsPoller.FinalResponse` has been removed
- Function `*JitRequestsCreateOrUpdatePoller.Poll` has been removed
- Function `*ApplicationDefinitionsListBySubscriptionPager.PageResponse` has been removed
- Function `*ApplicationsListBySubscriptionPager.Err` has been removed
- Function `PossibleManagedServiceIdentityTypeValues` has been removed
- Function `*ApplicationsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*ApplicationsListByResourceGroupPager.NextPage` has been removed
- Struct `ApplicationDefinitionsBeginCreateOrUpdateOptions` has been removed
- Struct `ApplicationDefinitionsBeginDeleteOptions` has been removed
- Struct `ApplicationDefinitionsCreateOrUpdatePoller` has been removed
- Struct `ApplicationDefinitionsCreateOrUpdatePollerResponse` has been removed
- Struct `ApplicationDefinitionsCreateOrUpdateResponse` has been removed
- Struct `ApplicationDefinitionsCreateOrUpdateResult` has been removed
- Struct `ApplicationDefinitionsDeletePoller` has been removed
- Struct `ApplicationDefinitionsDeletePollerResponse` has been removed
- Struct `ApplicationDefinitionsDeleteResponse` has been removed
- Struct `ApplicationDefinitionsGetOptions` has been removed
- Struct `ApplicationDefinitionsGetResponse` has been removed
- Struct `ApplicationDefinitionsGetResult` has been removed
- Struct `ApplicationDefinitionsListByResourceGroupOptions` has been removed
- Struct `ApplicationDefinitionsListByResourceGroupPager` has been removed
- Struct `ApplicationDefinitionsListByResourceGroupResponse` has been removed
- Struct `ApplicationDefinitionsListByResourceGroupResult` has been removed
- Struct `ApplicationDefinitionsListBySubscriptionOptions` has been removed
- Struct `ApplicationDefinitionsListBySubscriptionPager` has been removed
- Struct `ApplicationDefinitionsListBySubscriptionResponse` has been removed
- Struct `ApplicationDefinitionsListBySubscriptionResult` has been removed
- Struct `ApplicationDefinitionsUpdateOptions` has been removed
- Struct `ApplicationDefinitionsUpdateResponse` has been removed
- Struct `ApplicationDefinitionsUpdateResult` has been removed
- Struct `ApplicationsBeginCreateOrUpdateOptions` has been removed
- Struct `ApplicationsBeginDeleteOptions` has been removed
- Struct `ApplicationsBeginRefreshPermissionsOptions` has been removed
- Struct `ApplicationsCreateOrUpdatePoller` has been removed
- Struct `ApplicationsCreateOrUpdatePollerResponse` has been removed
- Struct `ApplicationsCreateOrUpdateResponse` has been removed
- Struct `ApplicationsCreateOrUpdateResult` has been removed
- Struct `ApplicationsDeletePoller` has been removed
- Struct `ApplicationsDeletePollerResponse` has been removed
- Struct `ApplicationsDeleteResponse` has been removed
- Struct `ApplicationsGetOptions` has been removed
- Struct `ApplicationsGetResponse` has been removed
- Struct `ApplicationsGetResult` has been removed
- Struct `ApplicationsListAllowedUpgradePlansOptions` has been removed
- Struct `ApplicationsListAllowedUpgradePlansResponse` has been removed
- Struct `ApplicationsListByResourceGroupOptions` has been removed
- Struct `ApplicationsListByResourceGroupPager` has been removed
- Struct `ApplicationsListByResourceGroupResponse` has been removed
- Struct `ApplicationsListByResourceGroupResult` has been removed
- Struct `ApplicationsListBySubscriptionOptions` has been removed
- Struct `ApplicationsListBySubscriptionPager` has been removed
- Struct `ApplicationsListBySubscriptionResponse` has been removed
- Struct `ApplicationsListBySubscriptionResult` has been removed
- Struct `ApplicationsRefreshPermissionsPoller` has been removed
- Struct `ApplicationsRefreshPermissionsPollerResponse` has been removed
- Struct `ApplicationsRefreshPermissionsResponse` has been removed
- Struct `ApplicationsUpdateOptions` has been removed
- Struct `ApplicationsUpdateResponse` has been removed
- Struct `ApplicationsUpdateResult` has been removed
- Struct `JitRequestsBeginCreateOrUpdateOptions` has been removed
- Struct `JitRequestsCreateOrUpdatePoller` has been removed
- Struct `JitRequestsCreateOrUpdatePollerResponse` has been removed
- Struct `JitRequestsCreateOrUpdateResponse` has been removed
- Struct `JitRequestsCreateOrUpdateResult` has been removed
- Struct `JitRequestsDeleteOptions` has been removed
- Struct `JitRequestsDeleteResponse` has been removed
- Struct `JitRequestsGetOptions` has been removed
- Struct `JitRequestsGetResponse` has been removed
- Struct `JitRequestsGetResult` has been removed
- Struct `JitRequestsListByResourceGroupOptions` has been removed
- Struct `JitRequestsListByResourceGroupResponse` has been removed
- Struct `JitRequestsListByResourceGroupResult` has been removed
- Struct `JitRequestsListBySubscriptionOptions` has been removed
- Struct `JitRequestsListBySubscriptionResponse` has been removed
- Struct `JitRequestsListBySubscriptionResult` has been removed
- Struct `JitRequestsUpdateOptions` has been removed
- Struct `JitRequestsUpdateResponse` has been removed
- Struct `JitRequestsUpdateResult` has been removed
- Struct `ManagedServiceIdentity` has been removed
- Struct `UserAssignedIdentity` has been removed
- Field `GenericResource` of struct `Application` has been removed
- Field `Resource` of struct `GenericResource` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed
- Field `GenericResource` of struct `ApplicationDefinition` has been removed
- Field `GenericResource` of struct `ApplicationPatchable` has been removed
- Field `Resource` of struct `JitRequestDefinition` has been removed

### Features Added

- New const `ResourceIdentityTypeNone`
- New const `ResourceIdentityTypeUserAssigned`
- New const `ResourceIdentityTypeSystemAssignedUserAssigned`
- New const `ResourceIdentityTypeSystemAssigned`
- New function `*JitRequestsClientCreateOrUpdatePollerResponse.Resume(context.Context, *JitRequestsClient, string) error`
- New function `*ApplicationsClientListBySubscriptionPager.Err() error`
- New function `*ApplicationsClientDeletePoller.Done() bool`
- New function `*ApplicationsClientRefreshPermissionsPoller.Poll(context.Context) (*http.Response, error)`
- New function `*ApplicationsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `ApplicationDefinitionsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ApplicationDefinitionsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationDefinitionsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ApplicationsClientListByResourceGroupPager.PageResponse() ApplicationsClientListByResourceGroupResponse`
- New function `*ApplicationsClientRefreshPermissionsPoller.FinalResponse(context.Context) (ApplicationsClientRefreshPermissionsResponse, error)`
- New function `*ApplicationDefinitionsClientListBySubscriptionPager.PageResponse() ApplicationDefinitionsClientListBySubscriptionResponse`
- New function `*JitRequestsClientCreateOrUpdatePoller.FinalResponse(context.Context) (JitRequestsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationsClientListBySubscriptionPager.PageResponse() ApplicationsClientListBySubscriptionResponse`
- New function `*ApplicationsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ApplicationsClientDeletePollerResponse.Resume(context.Context, *ApplicationsClient, string) error`
- New function `*ApplicationsClientRefreshPermissionsPollerResponse.Resume(context.Context, *ApplicationsClient, string) error`
- New function `*ApplicationsClientDeletePoller.ResumeToken() (string, error)`
- New function `*ApplicationDefinitionsClientCreateOrUpdatePollerResponse.Resume(context.Context, *ApplicationDefinitionsClient, string) error`
- New function `ApplicationsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (ApplicationsClientDeleteResponse, error)`
- New function `*JitRequestsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `ApplicationsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ApplicationsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationDefinitionsClientDeletePollerResponse.Resume(context.Context, *ApplicationDefinitionsClient, string) error`
- New function `*ApplicationDefinitionsClientListByResourceGroupPager.Err() error`
- New function `*ApplicationsClientListByResourceGroupPager.Err() error`
- New function `*ApplicationDefinitionsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ApplicationsClientDeletePoller.FinalResponse(context.Context) (ApplicationsClientDeleteResponse, error)`
- New function `*ApplicationDefinitionsClientListBySubscriptionPager.Err() error`
- New function `*ApplicationDefinitionsClientListByResourceGroupPager.PageResponse() ApplicationDefinitionsClientListByResourceGroupResponse`
- New function `JitRequestsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (JitRequestsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationDefinitionsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*ApplicationsClientCreateOrUpdatePoller.FinalResponse(context.Context) (ApplicationsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationDefinitionsClientCreateOrUpdatePoller.FinalResponse(context.Context) (ApplicationDefinitionsClientCreateOrUpdateResponse, error)`
- New function `ApplicationsClientRefreshPermissionsPollerResponse.PollUntilDone(context.Context, time.Duration) (ApplicationsClientRefreshPermissionsResponse, error)`
- New function `*ApplicationDefinitionsClientDeletePoller.ResumeToken() (string, error)`
- New function `*ApplicationsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*ApplicationsClientCreateOrUpdatePoller.Done() bool`
- New function `*ApplicationDefinitionsClientCreateOrUpdatePoller.Done() bool`
- New function `*ApplicationsClientRefreshPermissionsPoller.Done() bool`
- New function `Identity.MarshalJSON() ([]byte, error)`
- New function `*ApplicationDefinitionsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ApplicationsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ApplicationsClientRefreshPermissionsPoller.ResumeToken() (string, error)`
- New function `*JitRequestsClientCreateOrUpdatePoller.Done() bool`
- New function `*ApplicationsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*ApplicationDefinitionsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*JitRequestsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ApplicationDefinitionsClientDeletePoller.Done() bool`
- New function `*ApplicationDefinitionsClientDeletePoller.FinalResponse(context.Context) (ApplicationDefinitionsClientDeleteResponse, error)`
- New function `*ApplicationsClientCreateOrUpdatePollerResponse.Resume(context.Context, *ApplicationsClient, string) error`
- New function `PossibleResourceIdentityTypeValues() []ResourceIdentityType`
- New function `ApplicationDefinitionsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (ApplicationDefinitionsClientDeleteResponse, error)`
- New function `ResourceIdentityType.ToPtr() *ResourceIdentityType`
- New struct `ApplicationDefinitionsClientBeginCreateOrUpdateOptions`
- New struct `ApplicationDefinitionsClientBeginDeleteOptions`
- New struct `ApplicationDefinitionsClientCreateOrUpdatePoller`
- New struct `ApplicationDefinitionsClientCreateOrUpdatePollerResponse`
- New struct `ApplicationDefinitionsClientCreateOrUpdateResponse`
- New struct `ApplicationDefinitionsClientCreateOrUpdateResult`
- New struct `ApplicationDefinitionsClientDeletePoller`
- New struct `ApplicationDefinitionsClientDeletePollerResponse`
- New struct `ApplicationDefinitionsClientDeleteResponse`
- New struct `ApplicationDefinitionsClientGetOptions`
- New struct `ApplicationDefinitionsClientGetResponse`
- New struct `ApplicationDefinitionsClientGetResult`
- New struct `ApplicationDefinitionsClientListByResourceGroupOptions`
- New struct `ApplicationDefinitionsClientListByResourceGroupPager`
- New struct `ApplicationDefinitionsClientListByResourceGroupResponse`
- New struct `ApplicationDefinitionsClientListByResourceGroupResult`
- New struct `ApplicationDefinitionsClientListBySubscriptionOptions`
- New struct `ApplicationDefinitionsClientListBySubscriptionPager`
- New struct `ApplicationDefinitionsClientListBySubscriptionResponse`
- New struct `ApplicationDefinitionsClientListBySubscriptionResult`
- New struct `ApplicationDefinitionsClientUpdateOptions`
- New struct `ApplicationDefinitionsClientUpdateResponse`
- New struct `ApplicationDefinitionsClientUpdateResult`
- New struct `ApplicationsClientBeginCreateOrUpdateOptions`
- New struct `ApplicationsClientBeginDeleteOptions`
- New struct `ApplicationsClientBeginRefreshPermissionsOptions`
- New struct `ApplicationsClientCreateOrUpdatePoller`
- New struct `ApplicationsClientCreateOrUpdatePollerResponse`
- New struct `ApplicationsClientCreateOrUpdateResponse`
- New struct `ApplicationsClientCreateOrUpdateResult`
- New struct `ApplicationsClientDeletePoller`
- New struct `ApplicationsClientDeletePollerResponse`
- New struct `ApplicationsClientDeleteResponse`
- New struct `ApplicationsClientGetOptions`
- New struct `ApplicationsClientGetResponse`
- New struct `ApplicationsClientGetResult`
- New struct `ApplicationsClientListAllowedUpgradePlansOptions`
- New struct `ApplicationsClientListAllowedUpgradePlansResponse`
- New struct `ApplicationsClientListByResourceGroupOptions`
- New struct `ApplicationsClientListByResourceGroupPager`
- New struct `ApplicationsClientListByResourceGroupResponse`
- New struct `ApplicationsClientListByResourceGroupResult`
- New struct `ApplicationsClientListBySubscriptionOptions`
- New struct `ApplicationsClientListBySubscriptionPager`
- New struct `ApplicationsClientListBySubscriptionResponse`
- New struct `ApplicationsClientListBySubscriptionResult`
- New struct `ApplicationsClientRefreshPermissionsPoller`
- New struct `ApplicationsClientRefreshPermissionsPollerResponse`
- New struct `ApplicationsClientRefreshPermissionsResponse`
- New struct `ApplicationsClientUpdateOptions`
- New struct `ApplicationsClientUpdateResponse`
- New struct `ApplicationsClientUpdateResult`
- New struct `Identity`
- New struct `JitRequestsClientBeginCreateOrUpdateOptions`
- New struct `JitRequestsClientCreateOrUpdatePoller`
- New struct `JitRequestsClientCreateOrUpdatePollerResponse`
- New struct `JitRequestsClientCreateOrUpdateResponse`
- New struct `JitRequestsClientCreateOrUpdateResult`
- New struct `JitRequestsClientDeleteOptions`
- New struct `JitRequestsClientDeleteResponse`
- New struct `JitRequestsClientGetOptions`
- New struct `JitRequestsClientGetResponse`
- New struct `JitRequestsClientGetResult`
- New struct `JitRequestsClientListByResourceGroupOptions`
- New struct `JitRequestsClientListByResourceGroupResponse`
- New struct `JitRequestsClientListByResourceGroupResult`
- New struct `JitRequestsClientListBySubscriptionOptions`
- New struct `JitRequestsClientListBySubscriptionResponse`
- New struct `JitRequestsClientListBySubscriptionResult`
- New struct `JitRequestsClientUpdateOptions`
- New struct `JitRequestsClientUpdateResponse`
- New struct `JitRequestsClientUpdateResult`
- New struct `UserAssignedResourceIdentity`
- New field `ManagedBy` in struct `Application`
- New field `SKU` in struct `Application`
- New field `Tags` in struct `Application`
- New field `ID` in struct `Application`
- New field `Name` in struct `Application`
- New field `Type` in struct `Application`
- New field `Location` in struct `Application`
- New field `SystemData` in struct `Application`
- New field `SystemData` in struct `JitRequestDefinition`
- New field `Type` in struct `JitRequestDefinition`
- New field `Location` in struct `JitRequestDefinition`
- New field `Tags` in struct `JitRequestDefinition`
- New field `ID` in struct `JitRequestDefinition`
- New field `Name` in struct `JitRequestDefinition`
- New field `Name` in struct `ApplicationPatchable`
- New field `SystemData` in struct `ApplicationPatchable`
- New field `ID` in struct `ApplicationPatchable`
- New field `SKU` in struct `ApplicationPatchable`
- New field `Tags` in struct `ApplicationPatchable`
- New field `Type` in struct `ApplicationPatchable`
- New field `Location` in struct `ApplicationPatchable`
- New field `ManagedBy` in struct `ApplicationPatchable`
- New field `ID` in struct `GenericResource`
- New field `Name` in struct `GenericResource`
- New field `SystemData` in struct `GenericResource`
- New field `Type` in struct `GenericResource`
- New field `Location` in struct `GenericResource`
- New field `Tags` in struct `GenericResource`
- New field `Error` in struct `ErrorResponse`
- New field `Name` in struct `ApplicationDefinition`
- New field `ManagedBy` in struct `ApplicationDefinition`
- New field `Tags` in struct `ApplicationDefinition`
- New field `SystemData` in struct `ApplicationDefinition`
- New field `Type` in struct `ApplicationDefinition`
- New field `Location` in struct `ApplicationDefinition`
- New field `SKU` in struct `ApplicationDefinition`
- New field `ID` in struct `ApplicationDefinition`
- New field `AllowedDataActions` in struct `ApplicationPackageLockingPolicyDefinition`


## 0.1.0 (2021-12-23)

- Init release.