//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationAccountBatchConfigurationsClient contains the methods for the IntegrationAccountBatchConfigurations group.
// Don't use this type directly, use NewIntegrationAccountBatchConfigurationsClient() instead.
type IntegrationAccountBatchConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountBatchConfigurationsClient creates a new instance of IntegrationAccountBatchConfigurationsClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountBatchConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountBatchConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountBatchConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a batch configuration for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - batchConfigurationName - The batch configuration name.
//   - batchConfiguration - The batch configuration.
//   - options - IntegrationAccountBatchConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountBatchConfigurationsClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountBatchConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, batchConfiguration BatchConfiguration, options *IntegrationAccountBatchConfigurationsClientCreateOrUpdateOptions) (IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountBatchConfigurationsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, batchConfigurationName, batchConfiguration, options)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountBatchConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, batchConfiguration BatchConfiguration, options *IntegrationAccountBatchConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if batchConfigurationName == "" {
		return nil, errors.New("parameter batchConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{batchConfigurationName}", url.PathEscape(batchConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, batchConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountBatchConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchConfiguration); err != nil {
		return IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a batch configuration for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - batchConfigurationName - The batch configuration name.
//   - options - IntegrationAccountBatchConfigurationsClientDeleteOptions contains the optional parameters for the IntegrationAccountBatchConfigurationsClient.Delete
//     method.
func (client *IntegrationAccountBatchConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *IntegrationAccountBatchConfigurationsClientDeleteOptions) (IntegrationAccountBatchConfigurationsClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountBatchConfigurationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, batchConfigurationName, options)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountBatchConfigurationsClientDeleteResponse{}, err
	}
	return IntegrationAccountBatchConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountBatchConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *IntegrationAccountBatchConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if batchConfigurationName == "" {
		return nil, errors.New("parameter batchConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{batchConfigurationName}", url.PathEscape(batchConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a batch configuration for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - batchConfigurationName - The batch configuration name.
//   - options - IntegrationAccountBatchConfigurationsClientGetOptions contains the optional parameters for the IntegrationAccountBatchConfigurationsClient.Get
//     method.
func (client *IntegrationAccountBatchConfigurationsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *IntegrationAccountBatchConfigurationsClientGetOptions) (IntegrationAccountBatchConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountBatchConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, batchConfigurationName, options)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountBatchConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountBatchConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountBatchConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *IntegrationAccountBatchConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if batchConfigurationName == "" {
		return nil, errors.New("parameter batchConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{batchConfigurationName}", url.PathEscape(batchConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountBatchConfigurationsClient) getHandleResponse(resp *http.Response) (IntegrationAccountBatchConfigurationsClientGetResponse, error) {
	result := IntegrationAccountBatchConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchConfiguration); err != nil {
		return IntegrationAccountBatchConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the batch configurations for an integration account.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountBatchConfigurationsClientListOptions contains the optional parameters for the IntegrationAccountBatchConfigurationsClient.NewListPager
//     method.
func (client *IntegrationAccountBatchConfigurationsClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountBatchConfigurationsClientListOptions) *runtime.Pager[IntegrationAccountBatchConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountBatchConfigurationsClientListResponse]{
		More: func(page IntegrationAccountBatchConfigurationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountBatchConfigurationsClientListResponse) (IntegrationAccountBatchConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountBatchConfigurationsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			if err != nil {
				return IntegrationAccountBatchConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IntegrationAccountBatchConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationAccountBatchConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountBatchConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountBatchConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountBatchConfigurationsClient) listHandleResponse(resp *http.Response) (IntegrationAccountBatchConfigurationsClientListResponse, error) {
	result := IntegrationAccountBatchConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchConfigurationCollection); err != nil {
		return IntegrationAccountBatchConfigurationsClientListResponse{}, err
	}
	return result, nil
}