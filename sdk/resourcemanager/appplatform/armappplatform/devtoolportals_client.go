//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

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

// DevToolPortalsClient contains the methods for the DevToolPortals group.
// Don't use this type directly, use NewDevToolPortalsClient() instead.
type DevToolPortalsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDevToolPortalsClient creates a new instance of DevToolPortalsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDevToolPortalsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevToolPortalsClient, error) {
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
	client := &DevToolPortalsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default Dev Tool Portal or update the existing Dev Tool Portal.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// devToolPortalName - The name of Dev Tool Portal.
// devToolPortalResource - Parameters for the create or update operation
// options - DevToolPortalsClientBeginCreateOrUpdateOptions contains the optional parameters for the DevToolPortalsClient.BeginCreateOrUpdate
// method.
func (client *DevToolPortalsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, devToolPortalResource DevToolPortalResource, options *DevToolPortalsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevToolPortalsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, devToolPortalName, devToolPortalResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DevToolPortalsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DevToolPortalsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create the default Dev Tool Portal or update the existing Dev Tool Portal.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *DevToolPortalsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, devToolPortalResource DevToolPortalResource, options *DevToolPortalsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, devToolPortalName, devToolPortalResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DevToolPortalsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, devToolPortalResource DevToolPortalResource, options *DevToolPortalsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/DevToolPortals/{devToolPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if devToolPortalName == "" {
		return nil, errors.New("parameter devToolPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devToolPortalName}", url.PathEscape(devToolPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, devToolPortalResource)
}

// BeginDelete - Disable the default Dev Tool Portal.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// devToolPortalName - The name of Dev Tool Portal.
// options - DevToolPortalsClientBeginDeleteOptions contains the optional parameters for the DevToolPortalsClient.BeginDelete
// method.
func (client *DevToolPortalsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, options *DevToolPortalsClientBeginDeleteOptions) (*runtime.Poller[DevToolPortalsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, devToolPortalName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DevToolPortalsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DevToolPortalsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Disable the default Dev Tool Portal.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *DevToolPortalsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, options *DevToolPortalsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, devToolPortalName, options)
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
func (client *DevToolPortalsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, options *DevToolPortalsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/DevToolPortals/{devToolPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if devToolPortalName == "" {
		return nil, errors.New("parameter devToolPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devToolPortalName}", url.PathEscape(devToolPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Application Live and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// devToolPortalName - The name of Dev Tool Portal.
// options - DevToolPortalsClientGetOptions contains the optional parameters for the DevToolPortalsClient.Get method.
func (client *DevToolPortalsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, options *DevToolPortalsClientGetOptions) (DevToolPortalsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, devToolPortalName, options)
	if err != nil {
		return DevToolPortalsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevToolPortalsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevToolPortalsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevToolPortalsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, devToolPortalName string, options *DevToolPortalsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/DevToolPortals/{devToolPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if devToolPortalName == "" {
		return nil, errors.New("parameter devToolPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devToolPortalName}", url.PathEscape(devToolPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevToolPortalsClient) getHandleResponse(resp *http.Response) (DevToolPortalsClientGetResponse, error) {
	result := DevToolPortalsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevToolPortalResource); err != nil {
		return DevToolPortalsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - DevToolPortalsClientListOptions contains the optional parameters for the DevToolPortalsClient.List method.
func (client *DevToolPortalsClient) NewListPager(resourceGroupName string, serviceName string, options *DevToolPortalsClientListOptions) *runtime.Pager[DevToolPortalsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevToolPortalsClientListResponse]{
		More: func(page DevToolPortalsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevToolPortalsClientListResponse) (DevToolPortalsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevToolPortalsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevToolPortalsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevToolPortalsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DevToolPortalsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DevToolPortalsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/DevToolPortals"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DevToolPortalsClient) listHandleResponse(resp *http.Response) (DevToolPortalsClientListResponse, error) {
	result := DevToolPortalsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevToolPortalResourceCollection); err != nil {
		return DevToolPortalsClientListResponse{}, err
	}
	return result, nil
}
