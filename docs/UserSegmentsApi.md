# \UserSegmentsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserSegment**](UserSegmentsApi.md#DeleteUserSegment) | **Delete** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Delete a user segment.
[**GetExpiringUserTargetsOnSegment**](UserSegmentsApi.md#GetExpiringUserTargetsOnSegment) | **Get** /segments/{projectKey}/{userSegmentKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for user segment
[**GetUserSegment**](UserSegmentsApi.md#GetUserSegment) | **Get** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Get a single user segment by key.
[**GetUserSegments**](UserSegmentsApi.md#GetUserSegments) | **Get** /segments/{projectKey}/{environmentKey} | Get a list of all user segments in the given project.
[**PatchExpiringUserTargetsOnSegment**](UserSegmentsApi.md#PatchExpiringUserTargetsOnSegment) | **Patch** /segments/{projectKey}/{userSegmentKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets on user segment
[**PatchUserSegment**](UserSegmentsApi.md#PatchUserSegment) | **Patch** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Perform a partial update to a user segment.
[**PostUserSegment**](UserSegmentsApi.md#PostUserSegment) | **Post** /segments/{projectKey}/{environmentKey} | Creates a new user segment.
[**UpdatedUnboundedSegmentTargets**](UserSegmentsApi.md#UpdatedUnboundedSegmentTargets) | **Post** /segments/{projectKey}/{environmentKey}/{userSegmentKey}/unbounded-users | Update targets included or excluded in an unbounded segment


# **DeleteUserSegment**
> DeleteUserSegment(ctx, projectKey, environmentKey, userSegmentKey)
Delete a user segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiringUserTargetsOnSegment**
> UserTargetingExpirationForSegment GetExpiringUserTargetsOnSegment(ctx, projectKey, environmentKey, userSegmentKey)
Get expiring user targets for user segment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 

### Return type

[**UserTargetingExpirationForSegment**](UserTargetingExpirationForSegment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSegment**
> UserSegment GetUserSegment(ctx, projectKey, environmentKey, userSegmentKey)
Get a single user segment by key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSegments**
> UserSegments GetUserSegments(ctx, projectKey, environmentKey, optional)
Get a list of all user segments in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
 **optional** | ***UserSegmentsApiGetUserSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserSegmentsApiGetUserSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **optional.String**| Filter by tag. A tag can be used to group flags across projects. | 

### Return type

[**UserSegments**](UserSegments.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchExpiringUserTargetsOnSegment**
> UserTargetingExpirationForSegment PatchExpiringUserTargetsOnSegment(ctx, projectKey, environmentKey, userSegmentKey, semanticPatchWithComment)
Update, add, or delete expiring user targets on user segment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 
  **semanticPatchWithComment** | [**interface{}**](interface{}.md)| Requires a Semantic Patch representation of the desired changes to the resource. &#39;https://apidocs.launchdarkly.com/reference#updates-via-semantic-patches&#39;. The addition of comments is also supported. | 

### Return type

[**UserTargetingExpirationForSegment**](UserTargetingExpirationForSegment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUserSegment**
> UserSegment PatchUserSegment(ctx, projectKey, environmentKey, userSegmentKey, patchOnly)
Perform a partial update to a user segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 
  **patchOnly** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; Feature flag patches also support JSON Merge Patch format. &#39;https://tools.ietf.org/html/rfc7386&#39; The addition of comments is also supported. | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserSegment**
> UserSegment PostUserSegment(ctx, projectKey, environmentKey, userSegmentBody)
Creates a new user segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentBody** | [**UserSegmentBody**](UserSegmentBody.md)| Create a new user segment. | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatedUnboundedSegmentTargets**
> UpdatedUnboundedSegmentTargets(ctx, projectKey, environmentKey, userSegmentKey, unboundedSegmentTargetsBody)
Update targets included or excluded in an unbounded segment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userSegmentKey** | **string**| The user segment&#39;s key. The key identifies the user segment in your code. | 
  **unboundedSegmentTargetsBody** | [**UnboundedSegmentTargetsBody**](UnboundedSegmentTargetsBody.md)| Add or remove user targets to the included or excluded lists on an unbounded segment | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

