//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtimeseriesinsights

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

// EventSourcesClient contains the methods for the EventSources group.
// Don't use this type directly, use NewEventSourcesClient() instead.
type EventSourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEventSourcesClient creates a new instance of EventSourcesClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEventSourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EventSourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EventSourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an event source under the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-15
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - eventSourceName - Name of the event source.
//   - parameters - Parameters for creating an event source resource.
//   - options - EventSourcesClientCreateOrUpdateOptions contains the optional parameters for the EventSourcesClient.CreateOrUpdate
//     method.
func (client *EventSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, parameters EventSourceCreateOrUpdateParametersClassification, options *EventSourcesClientCreateOrUpdateOptions) (EventSourcesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "EventSourcesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, eventSourceName, parameters, options)
	if err != nil {
		return EventSourcesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventSourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EventSourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EventSourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, parameters EventSourceCreateOrUpdateParametersClassification, options *EventSourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if eventSourceName == "" {
		return nil, errors.New("parameter eventSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSourceName}", url.PathEscape(eventSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EventSourcesClient) createOrUpdateHandleResponse(resp *http.Response) (EventSourcesClientCreateOrUpdateResponse, error) {
	result := EventSourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EventSourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the event source with the specified name in the specified subscription, resource group, and environment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-15
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - eventSourceName - The name of the Time Series Insights event source associated with the specified environment.
//   - options - EventSourcesClientDeleteOptions contains the optional parameters for the EventSourcesClient.Delete method.
func (client *EventSourcesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, options *EventSourcesClientDeleteOptions) (EventSourcesClientDeleteResponse, error) {
	var err error
	const operationName = "EventSourcesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, eventSourceName, options)
	if err != nil {
		return EventSourcesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventSourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EventSourcesClientDeleteResponse{}, err
	}
	return EventSourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EventSourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, options *EventSourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if eventSourceName == "" {
		return nil, errors.New("parameter eventSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSourceName}", url.PathEscape(eventSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the event source with the specified name in the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-15
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - eventSourceName - The name of the Time Series Insights event source associated with the specified environment.
//   - options - EventSourcesClientGetOptions contains the optional parameters for the EventSourcesClient.Get method.
func (client *EventSourcesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, options *EventSourcesClientGetOptions) (EventSourcesClientGetResponse, error) {
	var err error
	const operationName = "EventSourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, eventSourceName, options)
	if err != nil {
		return EventSourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventSourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventSourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EventSourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, options *EventSourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if eventSourceName == "" {
		return nil, errors.New("parameter eventSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSourceName}", url.PathEscape(eventSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EventSourcesClient) getHandleResponse(resp *http.Response) (EventSourcesClientGetResponse, error) {
	result := EventSourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EventSourcesClientGetResponse{}, err
	}
	return result, nil
}

// ListByEnvironment - Lists all the available event sources associated with the subscription and within the specified resource
// group and environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-15
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - options - EventSourcesClientListByEnvironmentOptions contains the optional parameters for the EventSourcesClient.ListByEnvironment
//     method.
func (client *EventSourcesClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string, options *EventSourcesClientListByEnvironmentOptions) (EventSourcesClientListByEnvironmentResponse, error) {
	var err error
	const operationName = "EventSourcesClient.ListByEnvironment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByEnvironmentCreateRequest(ctx, resourceGroupName, environmentName, options)
	if err != nil {
		return EventSourcesClientListByEnvironmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventSourcesClientListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventSourcesClientListByEnvironmentResponse{}, err
	}
	resp, err := client.listByEnvironmentHandleResponse(httpResp)
	return resp, err
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *EventSourcesClient) listByEnvironmentCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *EventSourcesClientListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *EventSourcesClient) listByEnvironmentHandleResponse(resp *http.Response) (EventSourcesClientListByEnvironmentResponse, error) {
	result := EventSourcesClientListByEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSourceListResponse); err != nil {
		return EventSourcesClientListByEnvironmentResponse{}, err
	}
	return result, nil
}

// Update - Updates the event source with the specified name in the specified subscription, resource group, and environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-15
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - eventSourceName - The name of the Time Series Insights event source associated with the specified environment.
//   - eventSourceUpdateParameters - Request object that contains the updated information for the event source.
//   - options - EventSourcesClientUpdateOptions contains the optional parameters for the EventSourcesClient.Update method.
func (client *EventSourcesClient) Update(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, eventSourceUpdateParameters EventSourceUpdateParametersClassification, options *EventSourcesClientUpdateOptions) (EventSourcesClientUpdateResponse, error) {
	var err error
	const operationName = "EventSourcesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, environmentName, eventSourceName, eventSourceUpdateParameters, options)
	if err != nil {
		return EventSourcesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventSourcesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventSourcesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *EventSourcesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, eventSourceUpdateParameters EventSourceUpdateParametersClassification, options *EventSourcesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if eventSourceName == "" {
		return nil, errors.New("parameter eventSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSourceName}", url.PathEscape(eventSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSourceUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *EventSourcesClient) updateHandleResponse(resp *http.Response) (EventSourcesClientUpdateResponse, error) {
	result := EventSourcesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EventSourcesClientUpdateResponse{}, err
	}
	return result, nil
}