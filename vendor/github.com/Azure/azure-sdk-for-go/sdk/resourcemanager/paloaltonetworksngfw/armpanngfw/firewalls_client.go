//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallsClient contains the methods for the Firewalls group.
// Don't use this type directly, use NewFirewallsClient() instead.
type FirewallsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallsClient creates a new instance of FirewallsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - resource - Resource create parameters.
//   - options - FirewallsClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallsClient.BeginCreateOrUpdate
//     method.
func (client *FirewallsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallName string, resource FirewallResource, options *FirewallsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FirewallsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FirewallsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *FirewallsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallName string, resource FirewallResource, options *FirewallsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, resource FirewallResource, options *FirewallsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientBeginDeleteOptions contains the optional parameters for the FirewallsClient.BeginDelete method.
func (client *FirewallsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientBeginDeleteOptions) (*runtime.Poller[FirewallsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, firewallName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FirewallsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *FirewallsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientGetOptions contains the optional parameters for the FirewallsClient.Get method.
func (client *FirewallsClient) Get(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetOptions) (FirewallsClientGetResponse, error) {
	var err error
	const operationName = "FirewallsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return FirewallsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FirewallsClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallsClient) getHandleResponse(resp *http.Response) (FirewallsClientGetResponse, error) {
	result := FirewallsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallResource); err != nil {
		return FirewallsClientGetResponse{}, err
	}
	return result, nil
}

// GetGlobalRulestack - Get Global Rulestack associated with the Firewall
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientGetGlobalRulestackOptions contains the optional parameters for the FirewallsClient.GetGlobalRulestack
//     method.
func (client *FirewallsClient) GetGlobalRulestack(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetGlobalRulestackOptions) (FirewallsClientGetGlobalRulestackResponse, error) {
	var err error
	const operationName = "FirewallsClient.GetGlobalRulestack"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getGlobalRulestackCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return FirewallsClientGetGlobalRulestackResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientGetGlobalRulestackResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientGetGlobalRulestackResponse{}, err
	}
	resp, err := client.getGlobalRulestackHandleResponse(httpResp)
	return resp, err
}

// getGlobalRulestackCreateRequest creates the GetGlobalRulestack request.
func (client *FirewallsClient) getGlobalRulestackCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetGlobalRulestackOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}/getGlobalRulestack"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getGlobalRulestackHandleResponse handles the GetGlobalRulestack response.
func (client *FirewallsClient) getGlobalRulestackHandleResponse(resp *http.Response) (FirewallsClientGetGlobalRulestackResponse, error) {
	result := FirewallsClientGetGlobalRulestackResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalRulestackInfo); err != nil {
		return FirewallsClientGetGlobalRulestackResponse{}, err
	}
	return result, nil
}

// GetLogProfile - Log Profile for Firewall
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientGetLogProfileOptions contains the optional parameters for the FirewallsClient.GetLogProfile method.
func (client *FirewallsClient) GetLogProfile(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetLogProfileOptions) (FirewallsClientGetLogProfileResponse, error) {
	var err error
	const operationName = "FirewallsClient.GetLogProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getLogProfileCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return FirewallsClientGetLogProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientGetLogProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientGetLogProfileResponse{}, err
	}
	resp, err := client.getLogProfileHandleResponse(httpResp)
	return resp, err
}

// getLogProfileCreateRequest creates the GetLogProfile request.
func (client *FirewallsClient) getLogProfileCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetLogProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}/getLogProfile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogProfileHandleResponse handles the GetLogProfile response.
func (client *FirewallsClient) getLogProfileHandleResponse(resp *http.Response) (FirewallsClientGetLogProfileResponse, error) {
	result := FirewallsClientGetLogProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSettings); err != nil {
		return FirewallsClientGetLogProfileResponse{}, err
	}
	return result, nil
}

// GetSupportInfo - support info for firewall.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientGetSupportInfoOptions contains the optional parameters for the FirewallsClient.GetSupportInfo
//     method.
func (client *FirewallsClient) GetSupportInfo(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetSupportInfoOptions) (FirewallsClientGetSupportInfoResponse, error) {
	var err error
	const operationName = "FirewallsClient.GetSupportInfo"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSupportInfoCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return FirewallsClientGetSupportInfoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientGetSupportInfoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientGetSupportInfoResponse{}, err
	}
	resp, err := client.getSupportInfoHandleResponse(httpResp)
	return resp, err
}

// getSupportInfoCreateRequest creates the GetSupportInfo request.
func (client *FirewallsClient) getSupportInfoCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientGetSupportInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}/getSupportInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.Email != nil {
		reqQP.Set("email", *options.Email)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSupportInfoHandleResponse handles the GetSupportInfo response.
func (client *FirewallsClient) getSupportInfoHandleResponse(resp *http.Response) (FirewallsClientGetSupportInfoResponse, error) {
	result := FirewallsClientGetSupportInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SupportInfo); err != nil {
		return FirewallsClientGetSupportInfoResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List FirewallResource resources by resource group
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - FirewallsClientListByResourceGroupOptions contains the optional parameters for the FirewallsClient.NewListByResourceGroupPager
//     method.
func (client *FirewallsClient) NewListByResourceGroupPager(resourceGroupName string, options *FirewallsClientListByResourceGroupOptions) *runtime.Pager[FirewallsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallsClientListByResourceGroupResponse]{
		More: func(page FirewallsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallsClientListByResourceGroupResponse) (FirewallsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FirewallsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return FirewallsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FirewallsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FirewallsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FirewallsClient) listByResourceGroupHandleResponse(resp *http.Response) (FirewallsClientListByResourceGroupResponse, error) {
	result := FirewallsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallResourceListResult); err != nil {
		return FirewallsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List FirewallResource resources by subscription ID
//
// Generated from API version 2023-09-01
//   - options - FirewallsClientListBySubscriptionOptions contains the optional parameters for the FirewallsClient.NewListBySubscriptionPager
//     method.
func (client *FirewallsClient) NewListBySubscriptionPager(options *FirewallsClientListBySubscriptionOptions) *runtime.Pager[FirewallsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallsClientListBySubscriptionResponse]{
		More: func(page FirewallsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallsClientListBySubscriptionResponse) (FirewallsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FirewallsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return FirewallsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *FirewallsClient) listBySubscriptionCreateRequest(ctx context.Context, options *FirewallsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/PaloAltoNetworks.Cloudngfw/firewalls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *FirewallsClient) listBySubscriptionHandleResponse(resp *http.Response) (FirewallsClientListBySubscriptionResponse, error) {
	result := FirewallsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallResourceListResult); err != nil {
		return FirewallsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// SaveLogProfile - Log Profile for Firewall
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - options - FirewallsClientSaveLogProfileOptions contains the optional parameters for the FirewallsClient.SaveLogProfile
//     method.
func (client *FirewallsClient) SaveLogProfile(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientSaveLogProfileOptions) (FirewallsClientSaveLogProfileResponse, error) {
	var err error
	const operationName = "FirewallsClient.SaveLogProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.saveLogProfileCreateRequest(ctx, resourceGroupName, firewallName, options)
	if err != nil {
		return FirewallsClientSaveLogProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientSaveLogProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientSaveLogProfileResponse{}, err
	}
	return FirewallsClientSaveLogProfileResponse{}, nil
}

// saveLogProfileCreateRequest creates the SaveLogProfile request.
func (client *FirewallsClient) saveLogProfileCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, options *FirewallsClientSaveLogProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}/saveLogProfile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.LogSettings != nil {
		if err := runtime.MarshalAsJSON(req, *options.LogSettings); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Update - Update a FirewallResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - firewallName - Firewall resource name
//   - properties - The resource properties to be updated.
//   - options - FirewallsClientUpdateOptions contains the optional parameters for the FirewallsClient.Update method.
func (client *FirewallsClient) Update(ctx context.Context, resourceGroupName string, firewallName string, properties FirewallResourceUpdate, options *FirewallsClientUpdateOptions) (FirewallsClientUpdateResponse, error) {
	var err error
	const operationName = "FirewallsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, firewallName, properties, options)
	if err != nil {
		return FirewallsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *FirewallsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, firewallName string, properties FirewallResourceUpdate, options *FirewallsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/firewalls/{firewallName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallName == "" {
		return nil, errors.New("parameter firewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallName}", url.PathEscape(firewallName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *FirewallsClient) updateHandleResponse(resp *http.Response) (FirewallsClientUpdateResponse, error) {
	result := FirewallsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallResource); err != nil {
		return FirewallsClientUpdateResponse{}, err
	}
	return result, nil
}