package web

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// WorkflowRunActionRepetitionsClient is the webSite Management Client
type WorkflowRunActionRepetitionsClient struct {
	BaseClient
}

// NewWorkflowRunActionRepetitionsClient creates an instance of the WorkflowRunActionRepetitionsClient client.
func NewWorkflowRunActionRepetitionsClient(subscriptionID string) WorkflowRunActionRepetitionsClient {
	return NewWorkflowRunActionRepetitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunActionRepetitionsClientWithBaseURI creates an instance of the WorkflowRunActionRepetitionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewWorkflowRunActionRepetitionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunActionRepetitionsClient {
	return WorkflowRunActionRepetitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a workflow run action repetition.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - site name.
// workflowName - the workflow name.
// runName - the workflow run name.
// actionName - the workflow action name.
// repetitionName - the workflow repetition.
func (client WorkflowRunActionRepetitionsClient) Get(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string) (result WorkflowRunActionRepetitionDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRepetitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.WorkflowRunActionRepetitionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, name, workflowName, runName, actionName, repetitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunActionRepetitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        autorest.Encode("path", actionName),
		"name":              autorest.Encode("path", name),
		"repetitionName":    autorest.Encode("path", repetitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionRepetitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) GetResponder(resp *http.Response) (result WorkflowRunActionRepetitionDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all of a workflow run action repetitions.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - site name.
// workflowName - the workflow name.
// runName - the workflow run name.
// actionName - the workflow action name.
func (client WorkflowRunActionRepetitionsClient) List(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string) (result WorkflowRunActionRepetitionDefinitionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRepetitionsClient.List")
		defer func() {
			sc := -1
			if result.wrardc.Response.Response != nil {
				sc = result.wrardc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.WorkflowRunActionRepetitionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, name, workflowName, runName, actionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wrardc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.wrardc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.wrardc.hasNextLink() && result.wrardc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowRunActionRepetitionsClient) ListPreparer(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        autorest.Encode("path", actionName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionRepetitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) ListResponder(resp *http.Response) (result WorkflowRunActionRepetitionDefinitionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowRunActionRepetitionsClient) listNextResults(ctx context.Context, lastResults WorkflowRunActionRepetitionDefinitionCollection) (result WorkflowRunActionRepetitionDefinitionCollection, err error) {
	req, err := lastResults.workflowRunActionRepetitionDefinitionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowRunActionRepetitionsClient) ListComplete(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string) (result WorkflowRunActionRepetitionDefinitionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRepetitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, name, workflowName, runName, actionName)
	return
}

// ListExpressionTraces lists a workflow run expression trace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - site name.
// workflowName - the workflow name.
// runName - the workflow run name.
// actionName - the workflow action name.
// repetitionName - the workflow repetition.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTraces(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string) (result ExpressionTracesPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRepetitionsClient.ListExpressionTraces")
		defer func() {
			sc := -1
			if result.et.Response.Response != nil {
				sc = result.et.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", err.Error())
	}

	result.fn = client.listExpressionTracesNextResults
	req, err := client.ListExpressionTracesPreparer(ctx, resourceGroupName, name, workflowName, runName, actionName, repetitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListExpressionTracesSender(req)
	if err != nil {
		result.et.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", resp, "Failure sending request")
		return
	}

	result.et, err = client.ListExpressionTracesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", resp, "Failure responding to request")
		return
	}
	if result.et.hasNextLink() && result.et.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListExpressionTracesPreparer prepares the ListExpressionTraces request.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesPreparer(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        autorest.Encode("path", actionName),
		"name":              autorest.Encode("path", name),
		"repetitionName":    autorest.Encode("path", repetitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/listExpressionTraces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListExpressionTracesSender sends the ListExpressionTraces request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListExpressionTracesResponder handles the response to the ListExpressionTraces request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesResponder(resp *http.Response) (result ExpressionTraces, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listExpressionTracesNextResults retrieves the next set of results, if any.
func (client WorkflowRunActionRepetitionsClient) listExpressionTracesNextResults(ctx context.Context, lastResults ExpressionTraces) (result ExpressionTraces, err error) {
	req, err := lastResults.expressionTracesPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listExpressionTracesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListExpressionTracesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listExpressionTracesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListExpressionTracesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowRunActionRepetitionsClient", "listExpressionTracesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListExpressionTracesComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesComplete(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string) (result ExpressionTracesIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRepetitionsClient.ListExpressionTraces")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListExpressionTraces(ctx, resourceGroupName, name, workflowName, runName, actionName, repetitionName)
	return
}