//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CachesClient contains the methods for the Caches group.
// Don't use this type directly, use NewCachesClient() instead.
type CachesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCachesClient creates a new instance of CachesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCachesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CachesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CachesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Cache.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginCreateOrUpdateOptions contains the optional parameters for the CachesClient.BeginCreateOrUpdate
// method.
func (client *CachesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginCreateOrUpdateOptions) (*runtime.Poller[CachesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CachesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a Cache.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) createOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CachesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Cache != nil {
		return req, runtime.MarshalAsJSON(req, *options.Cache)
	}
	return req, nil
}

// BeginDebugInfo - Tells a Cache to write generate debug info for support to process.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginDebugInfoOptions contains the optional parameters for the CachesClient.BeginDebugInfo method.
func (client *CachesClient) BeginDebugInfo(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDebugInfoOptions) (*runtime.Poller[CachesClientDebugInfoResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.debugInfo(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CachesClientDebugInfoResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientDebugInfoResponse](options.ResumeToken, client.pl, nil)
	}
}

// DebugInfo - Tells a Cache to write generate debug info for support to process.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) debugInfo(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDebugInfoOptions) (*http.Response, error) {
	req, err := client.debugInfoCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// debugInfoCreateRequest creates the DebugInfo request.
func (client *CachesClient) debugInfoCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDebugInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/debugInfo"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Schedules a Cache for deletion.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginDeleteOptions contains the optional parameters for the CachesClient.BeginDelete method.
func (client *CachesClient) BeginDelete(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDeleteOptions) (*runtime.Poller[CachesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CachesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Schedules a Cache for deletion.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) deleteOperation(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CachesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginFlush - Tells a Cache to write all dirty data to the Storage Target(s). During the flush, clients will see errors
// returned until the flush is complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginFlushOptions contains the optional parameters for the CachesClient.BeginFlush method.
func (client *CachesClient) BeginFlush(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginFlushOptions) (*runtime.Poller[CachesClientFlushResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.flush(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CachesClientFlushResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientFlushResponse](options.ResumeToken, client.pl, nil)
	}
}

// Flush - Tells a Cache to write all dirty data to the Storage Target(s). During the flush, clients will see errors returned
// until the flush is complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) flush(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginFlushOptions) (*http.Response, error) {
	req, err := client.flushCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// flushCreateRequest creates the Flush request.
func (client *CachesClient) flushCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginFlushOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/flush"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a Cache.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientGetOptions contains the optional parameters for the CachesClient.Get method.
func (client *CachesClient) Get(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientGetOptions) (CachesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return CachesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CachesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CachesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CachesClient) getCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CachesClient) getHandleResponse(resp *http.Response) (CachesClientGetResponse, error) {
	result := CachesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cache); err != nil {
		return CachesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns all Caches the user has access to under a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// options - CachesClientListOptions contains the optional parameters for the CachesClient.List method.
func (client *CachesClient) NewListPager(options *CachesClientListOptions) *runtime.Pager[CachesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CachesClientListResponse]{
		More: func(page CachesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CachesClientListResponse) (CachesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CachesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CachesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CachesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CachesClient) listCreateRequest(ctx context.Context, options *CachesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CachesClient) listHandleResponse(resp *http.Response) (CachesClientListResponse, error) {
	result := CachesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CachesListResult); err != nil {
		return CachesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns all Caches the user has access to under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// options - CachesClientListByResourceGroupOptions contains the optional parameters for the CachesClient.ListByResourceGroup
// method.
func (client *CachesClient) NewListByResourceGroupPager(resourceGroupName string, options *CachesClientListByResourceGroupOptions) *runtime.Pager[CachesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CachesClientListByResourceGroupResponse]{
		More: func(page CachesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CachesClientListByResourceGroupResponse) (CachesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CachesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CachesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CachesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CachesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CachesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CachesClient) listByResourceGroupHandleResponse(resp *http.Response) (CachesClientListByResourceGroupResponse, error) {
	result := CachesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CachesListResult); err != nil {
		return CachesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginStart - Tells a Stopped state Cache to transition to Active state.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginStartOptions contains the optional parameters for the CachesClient.BeginStart method.
func (client *CachesClient) BeginStart(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStartOptions) (*runtime.Poller[CachesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CachesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientStartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Start - Tells a Stopped state Cache to transition to Active state.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) start(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *CachesClient) startCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Tells an Active Cache to transition to Stopped state.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginStopOptions contains the optional parameters for the CachesClient.BeginStop method.
func (client *CachesClient) BeginStop(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStopOptions) (*runtime.Poller[CachesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CachesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientStopResponse](options.ResumeToken, client.pl, nil)
	}
}

// Stop - Tells an Active Cache to transition to Stopped state.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) stop(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *CachesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update a Cache instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientUpdateOptions contains the optional parameters for the CachesClient.Update method.
func (client *CachesClient) Update(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientUpdateOptions) (CachesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return CachesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CachesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CachesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CachesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Cache != nil {
		return req, runtime.MarshalAsJSON(req, *options.Cache)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CachesClient) updateHandleResponse(resp *http.Response) (CachesClientUpdateResponse, error) {
	result := CachesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cache); err != nil {
		return CachesClientUpdateResponse{}, err
	}
	return result, nil
}

// BeginUpgradeFirmware - Upgrade a Cache's firmware if a new version is available. Otherwise, this operation has no effect.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - CachesClientBeginUpgradeFirmwareOptions contains the optional parameters for the CachesClient.BeginUpgradeFirmware
// method.
func (client *CachesClient) BeginUpgradeFirmware(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginUpgradeFirmwareOptions) (*runtime.Poller[CachesClientUpgradeFirmwareResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgradeFirmware(ctx, resourceGroupName, cacheName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CachesClientUpgradeFirmwareResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CachesClientUpgradeFirmwareResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpgradeFirmware - Upgrade a Cache's firmware if a new version is available. Otherwise, this operation has no effect.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *CachesClient) upgradeFirmware(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginUpgradeFirmwareOptions) (*http.Response, error) {
	req, err := client.upgradeFirmwareCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeFirmwareCreateRequest creates the UpgradeFirmware request.
func (client *CachesClient) upgradeFirmwareCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *CachesClientBeginUpgradeFirmwareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/upgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}