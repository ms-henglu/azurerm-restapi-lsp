//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolDataWarehouseUserActivitiesClient contains the methods for the SQLPoolDataWarehouseUserActivities group.
// Don't use this type directly, use NewSQLPoolDataWarehouseUserActivitiesClient() instead.
type SQLPoolDataWarehouseUserActivitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLPoolDataWarehouseUserActivitiesClient creates a new instance of SQLPoolDataWarehouseUserActivitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolDataWarehouseUserActivitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolDataWarehouseUserActivitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLPoolDataWarehouseUserActivitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the user activities of a SQL pool which includes running and suspended queries
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - dataWarehouseUserActivityName - The activity name of the Sql pool.
//   - options - SQLPoolDataWarehouseUserActivitiesClientGetOptions contains the optional parameters for the SQLPoolDataWarehouseUserActivitiesClient.Get
//     method.
func (client *SQLPoolDataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *SQLPoolDataWarehouseUserActivitiesClientGetOptions) (SQLPoolDataWarehouseUserActivitiesClientGetResponse, error) {
	var err error
	const operationName = "SQLPoolDataWarehouseUserActivitiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, dataWarehouseUserActivityName, options)
	if err != nil {
		return SQLPoolDataWarehouseUserActivitiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLPoolDataWarehouseUserActivitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLPoolDataWarehouseUserActivitiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLPoolDataWarehouseUserActivitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *SQLPoolDataWarehouseUserActivitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if dataWarehouseUserActivityName == "" {
		return nil, errors.New("parameter dataWarehouseUserActivityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataWarehouseUserActivityName}", url.PathEscape(string(dataWarehouseUserActivityName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolDataWarehouseUserActivitiesClient) getHandleResponse(resp *http.Response) (SQLPoolDataWarehouseUserActivitiesClientGetResponse, error) {
	result := SQLPoolDataWarehouseUserActivitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataWarehouseUserActivities); err != nil {
		return SQLPoolDataWarehouseUserActivitiesClientGetResponse{}, err
	}
	return result, nil
}