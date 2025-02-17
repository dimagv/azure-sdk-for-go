//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

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

// WorkflowsClient contains the methods for the Workflows group.
// Don't use this type directly, use NewWorkflowsClient() instead.
type WorkflowsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowsClient creates a new instance of WorkflowsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowsClient, error) {
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
	client := &WorkflowsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Abort - Abort the given workflow.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// workflowID - workflow Id
// options - WorkflowsClientAbortOptions contains the optional parameters for the WorkflowsClient.Abort method.
func (client *WorkflowsClient) Abort(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *WorkflowsClientAbortOptions) (WorkflowsClientAbortResponse, error) {
	req, err := client.abortCreateRequest(ctx, resourceGroupName, storageSyncServiceName, workflowID, options)
	if err != nil {
		return WorkflowsClientAbortResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowsClientAbortResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowsClientAbortResponse{}, runtime.NewResponseError(resp)
	}
	return client.abortHandleResponse(resp)
}

// abortCreateRequest creates the Abort request.
func (client *WorkflowsClient) abortCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *WorkflowsClientAbortOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows/{workflowId}/abort"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if workflowID == "" {
		return nil, errors.New("parameter workflowID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowId}", url.PathEscape(workflowID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// abortHandleResponse handles the Abort response.
func (client *WorkflowsClient) abortHandleResponse(resp *http.Response) (WorkflowsClientAbortResponse, error) {
	result := WorkflowsClientAbortResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	return result, nil
}

// Get - Get Workflows resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// workflowID - workflow Id
// options - WorkflowsClientGetOptions contains the optional parameters for the WorkflowsClient.Get method.
func (client *WorkflowsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *WorkflowsClientGetOptions) (WorkflowsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, workflowID, options)
	if err != nil {
		return WorkflowsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *WorkflowsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows/{workflowId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if workflowID == "" {
		return nil, errors.New("parameter workflowID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowId}", url.PathEscape(workflowID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowsClient) getHandleResponse(resp *http.Response) (WorkflowsClientGetResponse, error) {
	result := WorkflowsClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workflow); err != nil {
		return WorkflowsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByStorageSyncServicePager - Get a Workflow List
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// options - WorkflowsClientListByStorageSyncServiceOptions contains the optional parameters for the WorkflowsClient.ListByStorageSyncService
// method.
func (client *WorkflowsClient) NewListByStorageSyncServicePager(resourceGroupName string, storageSyncServiceName string, options *WorkflowsClientListByStorageSyncServiceOptions) *runtime.Pager[WorkflowsClientListByStorageSyncServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowsClientListByStorageSyncServiceResponse]{
		More: func(page WorkflowsClientListByStorageSyncServiceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkflowsClientListByStorageSyncServiceResponse) (WorkflowsClientListByStorageSyncServiceResponse, error) {
			req, err := client.listByStorageSyncServiceCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
			if err != nil {
				return WorkflowsClientListByStorageSyncServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowsClientListByStorageSyncServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowsClientListByStorageSyncServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByStorageSyncServiceHandleResponse(resp)
		},
	})
}

// listByStorageSyncServiceCreateRequest creates the ListByStorageSyncService request.
func (client *WorkflowsClient) listByStorageSyncServiceCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *WorkflowsClientListByStorageSyncServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStorageSyncServiceHandleResponse handles the ListByStorageSyncService response.
func (client *WorkflowsClient) listByStorageSyncServiceHandleResponse(resp *http.Response) (WorkflowsClientListByStorageSyncServiceResponse, error) {
	result := WorkflowsClientListByStorageSyncServiceResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowArray); err != nil {
		return WorkflowsClientListByStorageSyncServiceResponse{}, err
	}
	return result, nil
}
