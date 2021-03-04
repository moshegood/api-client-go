This repository contains a client library for LaunchDarkly's REST API. This client was automatically
generated from our [OpenAPI specification](https://github.com/launchdarkly/ld-openapi).

This REST API is for custom integrations, data export, or automating your feature flag workflows. *DO NOT* use this client library to include feature flags in your web or mobile application. To integrate feature flags with your application, please see the [SDK documentation](https://docs.launchdarkly.com/v2.0/docs)

# Go API client for ldapi

Build custom integrations with the LaunchDarkly REST API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 5.0.2
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://support.launchdarkly.com](https://support.launchdarkly.com)

## Installation

```golang
import "github.com/launchdarkly/api-client-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessTokensApi* | [**DeleteToken**](docs/AccessTokensApi.md#deletetoken) | **Delete** /tokens/{tokenId} | Delete an access token by ID.
*AccessTokensApi* | [**GetToken**](docs/AccessTokensApi.md#gettoken) | **Get** /tokens/{tokenId} | Get a single access token by ID.
*AccessTokensApi* | [**GetTokens**](docs/AccessTokensApi.md#gettokens) | **Get** /tokens | Returns a list of tokens in the account.
*AccessTokensApi* | [**PatchToken**](docs/AccessTokensApi.md#patchtoken) | **Patch** /tokens/{tokenId} | Modify an access token by ID.
*AccessTokensApi* | [**PostToken**](docs/AccessTokensApi.md#posttoken) | **Post** /tokens | Create a new token.
*AccessTokensApi* | [**ResetToken**](docs/AccessTokensApi.md#resettoken) | **Post** /tokens/{tokenId}/reset | Reset an access token&#39;s secret key with an optional expiry time for the old key.
*AuditLogApi* | [**GetAuditLogEntries**](docs/AuditLogApi.md#getauditlogentries) | **Get** /auditlog | Get a list of all audit log entries. The query parameters allow you to restrict the returned results by date ranges, resource specifiers, or a full-text search query.
*AuditLogApi* | [**GetAuditLogEntry**](docs/AuditLogApi.md#getauditlogentry) | **Get** /auditlog/{resourceId} | Use this endpoint to fetch a single audit log entry by its resouce ID.
*CustomRolesApi* | [**DeleteCustomRole**](docs/CustomRolesApi.md#deletecustomrole) | **Delete** /roles/{customRoleKey} | Delete a custom role by key.
*CustomRolesApi* | [**GetCustomRole**](docs/CustomRolesApi.md#getcustomrole) | **Get** /roles/{customRoleKey} | Get one custom role by key.
*CustomRolesApi* | [**GetCustomRoles**](docs/CustomRolesApi.md#getcustomroles) | **Get** /roles | Return a complete list of custom roles.
*CustomRolesApi* | [**PatchCustomRole**](docs/CustomRolesApi.md#patchcustomrole) | **Patch** /roles/{customRoleKey} | Modify a custom role by key.
*CustomRolesApi* | [**PostCustomRole**](docs/CustomRolesApi.md#postcustomrole) | **Post** /roles | Create a new custom role.
*CustomerMetricsApi* | [**GetEvaluations**](docs/CustomerMetricsApi.md#getevaluations) | **Get** /usage/evaluations/{envId}/{flagKey} | Get events usage by event id and the feature flag key.
*CustomerMetricsApi* | [**GetEvent**](docs/CustomerMetricsApi.md#getevent) | **Get** /usage/events/{type} | Get events usage by event type.
*CustomerMetricsApi* | [**GetEvents**](docs/CustomerMetricsApi.md#getevents) | **Get** /usage/events | Get events usage endpoints.
*CustomerMetricsApi* | [**GetMAU**](docs/CustomerMetricsApi.md#getmau) | **Get** /usage/mau | Get monthly active user data.
*CustomerMetricsApi* | [**GetMAUByCategory**](docs/CustomerMetricsApi.md#getmaubycategory) | **Get** /usage/mau/bycategory | Get monthly active user data by category.
*CustomerMetricsApi* | [**GetStream**](docs/CustomerMetricsApi.md#getstream) | **Get** /usage/streams/{source} | Get a stream endpoint and return timeseries data.
*CustomerMetricsApi* | [**GetStreamBySDK**](docs/CustomerMetricsApi.md#getstreambysdk) | **Get** /usage/streams/{source}/bysdkversion | Get a stream timeseries data by source show sdk version metadata.
*CustomerMetricsApi* | [**GetStreamSDKVersion**](docs/CustomerMetricsApi.md#getstreamsdkversion) | **Get** /usage/streams/{source}/sdkversions | Get a stream timeseries data by source and show all sdk version associated.
*CustomerMetricsApi* | [**GetStreams**](docs/CustomerMetricsApi.md#getstreams) | **Get** /usage/streams | Returns a list of all streams.
*CustomerMetricsApi* | [**GetUsage**](docs/CustomerMetricsApi.md#getusage) | **Get** /usage | Returns of the usage endpoints available.
*DataExportDestinationsApi* | [**DeleteDestination**](docs/DataExportDestinationsApi.md#deletedestination) | **Delete** /destinations/{projectKey}/{environmentKey}/{destinationId} | Get a single data export destination by ID
*DataExportDestinationsApi* | [**GetDestination**](docs/DataExportDestinationsApi.md#getdestination) | **Get** /destinations/{projectKey}/{environmentKey}/{destinationId} | Get a single data export destination by ID
*DataExportDestinationsApi* | [**GetDestinations**](docs/DataExportDestinationsApi.md#getdestinations) | **Get** /destinations | Returns a list of all data export destinations.
*DataExportDestinationsApi* | [**PatchDestination**](docs/DataExportDestinationsApi.md#patchdestination) | **Patch** /destinations/{projectKey}/{environmentKey}/{destinationId} | Perform a partial update to a data export destination.
*DataExportDestinationsApi* | [**PostDestination**](docs/DataExportDestinationsApi.md#postdestination) | **Post** /destinations/{projectKey}/{environmentKey} | Create a new data export destination
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /projects/{projectKey}/environments/{environmentKey} | Delete an environment in a specific project.
*EnvironmentsApi* | [**GetEnvironment**](docs/EnvironmentsApi.md#getenvironment) | **Get** /projects/{projectKey}/environments/{environmentKey} | Get an environment given a project and key.
*EnvironmentsApi* | [**PatchEnvironment**](docs/EnvironmentsApi.md#patchenvironment) | **Patch** /projects/{projectKey}/environments/{environmentKey} | Modify an environment by ID.
*EnvironmentsApi* | [**PostEnvironment**](docs/EnvironmentsApi.md#postenvironment) | **Post** /projects/{projectKey}/environments | Create a new environment in a specified project with a given name, key, and swatch color.
*EnvironmentsApi* | [**ResetEnvironmentMobileKey**](docs/EnvironmentsApi.md#resetenvironmentmobilekey) | **Post** /projects/{projectKey}/environments/{environmentKey}/mobileKey | Reset an environment&#39;s mobile key. The optional expiry for the old key is deprecated for this endpoint, so the old key will always expire immediately.
*EnvironmentsApi* | [**ResetEnvironmentSDKKey**](docs/EnvironmentsApi.md#resetenvironmentsdkkey) | **Post** /projects/{projectKey}/environments/{environmentKey}/apiKey | Reset an environment&#39;s SDK key with an optional expiry time for the old key.
*FeatureFlagsApi* | [**CopyFeatureFlag**](docs/FeatureFlagsApi.md#copyfeatureflag) | **Post** /flags/{projectKey}/{featureFlagKey}/copy | Copies the feature flag configuration from one environment to the same feature flag in another environment.
*FeatureFlagsApi* | [**DeleteApprovalRequest**](docs/FeatureFlagsApi.md#deleteapprovalrequest) | **Delete** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Delete an approval request for a feature flag config
*FeatureFlagsApi* | [**DeleteFeatureFlag**](docs/FeatureFlagsApi.md#deletefeatureflag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.
*FeatureFlagsApi* | [**DeleteFlagConfigScheduledChanges**](docs/FeatureFlagsApi.md#deleteflagconfigscheduledchanges) | **Delete** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Delete a scheduled change on a feature flag in an environment.
*FeatureFlagsApi* | [**GetApprovalRequest**](docs/FeatureFlagsApi.md#getapprovalrequest) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Get a single approval request for a feature flag config
*FeatureFlagsApi* | [**GetApprovalRequests**](docs/FeatureFlagsApi.md#getapprovalrequests) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Get all approval requests for a feature flag config
*FeatureFlagsApi* | [**GetExpiringUserTargets**](docs/FeatureFlagsApi.md#getexpiringusertargets) | **Get** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for feature flag
*FeatureFlagsApi* | [**GetFeatureFlag**](docs/FeatureFlagsApi.md#getfeatureflag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
*FeatureFlagsApi* | [**GetFeatureFlagStatus**](docs/FeatureFlagsApi.md#getfeatureflagstatus) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get the status for a particular feature flag.
*FeatureFlagsApi* | [**GetFeatureFlagStatusAcrossEnvironments**](docs/FeatureFlagsApi.md#getfeatureflagstatusacrossenvironments) | **Get** /flag-status/{projectKey}/{featureFlagKey} | Get the status for a particular feature flag across environments
*FeatureFlagsApi* | [**GetFeatureFlagStatuses**](docs/FeatureFlagsApi.md#getfeatureflagstatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.
*FeatureFlagsApi* | [**GetFeatureFlags**](docs/FeatureFlagsApi.md#getfeatureflags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
*FeatureFlagsApi* | [**GetFlagConfigScheduledChange**](docs/FeatureFlagsApi.md#getflagconfigscheduledchange) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Get a scheduled change on a feature flag by id.
*FeatureFlagsApi* | [**GetFlagConfigScheduledChanges**](docs/FeatureFlagsApi.md#getflagconfigscheduledchanges) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Get all scheduled workflows for a feature flag by key.
*FeatureFlagsApi* | [**GetFlagConfigScheduledChangesConflicts**](docs/FeatureFlagsApi.md#getflagconfigscheduledchangesconflicts) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes-conflicts | Lists conflicts between the given instructions and any existing scheduled changes for the feature flag. The actual HTTP verb should be REPORT, not POST.
*FeatureFlagsApi* | [**PatchExpiringUserTargets**](docs/FeatureFlagsApi.md#patchexpiringusertargets) | **Patch** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets on feature flag
*FeatureFlagsApi* | [**PatchFeatureFlag**](docs/FeatureFlagsApi.md#patchfeatureflag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Perform a partial update to a feature.
*FeatureFlagsApi* | [**PatchFlagConfigScheduledChange**](docs/FeatureFlagsApi.md#patchflagconfigscheduledchange) | **Patch** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Updates an existing scheduled-change on a feature flag in an environment.
*FeatureFlagsApi* | [**PostApplyApprovalRequest**](docs/FeatureFlagsApi.md#postapplyapprovalrequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId}/apply | Apply approval request for a feature flag config
*FeatureFlagsApi* | [**PostApprovalRequest**](docs/FeatureFlagsApi.md#postapprovalrequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Create an approval request for a feature flag config
*FeatureFlagsApi* | [**PostFeatureFlag**](docs/FeatureFlagsApi.md#postfeatureflag) | **Post** /flags/{projectKey} | Creates a new feature flag.
*FeatureFlagsApi* | [**PostFlagConfigScheduledChanges**](docs/FeatureFlagsApi.md#postflagconfigscheduledchanges) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Creates a new scheduled change for a feature flag.
*FeatureFlagsApi* | [**PostReviewApprovalRequest**](docs/FeatureFlagsApi.md#postreviewapprovalrequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId}/review | Review approval request for a feature flag config
*IntegrationsApi* | [**DeleteIntegrationSubscription**](docs/IntegrationsApi.md#deleteintegrationsubscription) | **Delete** /integrations/{integrationKey}/{integrationId} | Delete an integration subscription by ID.
*IntegrationsApi* | [**GetIntegrationSubscription**](docs/IntegrationsApi.md#getintegrationsubscription) | **Get** /integrations/{integrationKey}/{integrationId} | Get a single integration subscription by ID.
*IntegrationsApi* | [**GetIntegrationSubscriptions**](docs/IntegrationsApi.md#getintegrationsubscriptions) | **Get** /integrations/{integrationKey} | Get a list of all configured integrations of a given kind.
*IntegrationsApi* | [**GetIntegrations**](docs/IntegrationsApi.md#getintegrations) | **Get** /integrations | Get a list of all configured audit log event integrations associated with this account.
*IntegrationsApi* | [**PatchIntegrationSubscription**](docs/IntegrationsApi.md#patchintegrationsubscription) | **Patch** /integrations/{integrationKey}/{integrationId} | Modify an integration subscription by ID.
*IntegrationsApi* | [**PostIntegrationSubscription**](docs/IntegrationsApi.md#postintegrationsubscription) | **Post** /integrations/{integrationKey} | Create a new integration subscription of a given kind.
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /projects/{projectKey} | Delete a project by key. Caution-- deleting a project will delete all associated environments and feature flags. You cannot delete the last project in an account.
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /projects/{projectKey} | Fetch a single project by key.
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /projects | Returns a list of all projects in the account.
*ProjectsApi* | [**PatchProject**](docs/ProjectsApi.md#patchproject) | **Patch** /projects/{projectKey} | Modify a project by ID.
*ProjectsApi* | [**PostProject**](docs/ProjectsApi.md#postproject) | **Post** /projects | Create a new project with the given key and name.
*RelayProxyConfigurationsApi* | [**DeleteRelayProxyConfig**](docs/RelayProxyConfigurationsApi.md#deleterelayproxyconfig) | **Delete** /account/relay-auto-configs/{id} | Delete a relay proxy configuration by ID.
*RelayProxyConfigurationsApi* | [**GetRelayProxyConfig**](docs/RelayProxyConfigurationsApi.md#getrelayproxyconfig) | **Get** /account/relay-auto-configs/{id} | Get a single relay proxy configuration by ID.
*RelayProxyConfigurationsApi* | [**GetRelayProxyConfigs**](docs/RelayProxyConfigurationsApi.md#getrelayproxyconfigs) | **Get** /account/relay-auto-configs | Returns a list of relay proxy configurations in the account.
*RelayProxyConfigurationsApi* | [**PatchRelayProxyConfig**](docs/RelayProxyConfigurationsApi.md#patchrelayproxyconfig) | **Patch** /account/relay-auto-configs/{id} | Modify a relay proxy configuration by ID.
*RelayProxyConfigurationsApi* | [**PostRelayAutoConfig**](docs/RelayProxyConfigurationsApi.md#postrelayautoconfig) | **Post** /account/relay-auto-configs | Create a new relay proxy config.
*RelayProxyConfigurationsApi* | [**ResetRelayProxyConfig**](docs/RelayProxyConfigurationsApi.md#resetrelayproxyconfig) | **Post** /account/relay-auto-configs/{id}/reset | Reset a relay proxy configuration&#39;s secret key with an optional expiry time for the old key.
*RootApi* | [**GetRoot**](docs/RootApi.md#getroot) | **Get** / | 
*TeamMembersApi* | [**DeleteMember**](docs/TeamMembersApi.md#deletemember) | **Delete** /members/{memberId} | Delete a team member by ID.
*TeamMembersApi* | [**GetMe**](docs/TeamMembersApi.md#getme) | **Get** /members/me | Get the current team member associated with the token
*TeamMembersApi* | [**GetMember**](docs/TeamMembersApi.md#getmember) | **Get** /members/{memberId} | Get a single team member by ID.
*TeamMembersApi* | [**GetMembers**](docs/TeamMembersApi.md#getmembers) | **Get** /members | Returns a list of all members in the account.
*TeamMembersApi* | [**PatchMember**](docs/TeamMembersApi.md#patchmember) | **Patch** /members/{memberId} | Modify a team member by ID.
*TeamMembersApi* | [**PostMembers**](docs/TeamMembersApi.md#postmembers) | **Post** /members | Invite new members.
*UserSegmentsApi* | [**DeleteUserSegment**](docs/UserSegmentsApi.md#deleteusersegment) | **Delete** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Delete a user segment.
*UserSegmentsApi* | [**GetExpiringUserTargetsOnSegment**](docs/UserSegmentsApi.md#getexpiringusertargetsonsegment) | **Get** /segments/{projectKey}/{userSegmentKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for user segment
*UserSegmentsApi* | [**GetUserSegment**](docs/UserSegmentsApi.md#getusersegment) | **Get** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Get a single user segment by key.
*UserSegmentsApi* | [**GetUserSegments**](docs/UserSegmentsApi.md#getusersegments) | **Get** /segments/{projectKey}/{environmentKey} | Get a list of all user segments in the given project.
*UserSegmentsApi* | [**PatchExpiringUserTargetsOnSegment**](docs/UserSegmentsApi.md#patchexpiringusertargetsonsegment) | **Patch** /segments/{projectKey}/{userSegmentKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets on user segment
*UserSegmentsApi* | [**PatchUserSegment**](docs/UserSegmentsApi.md#patchusersegment) | **Patch** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Perform a partial update to a user segment.
*UserSegmentsApi* | [**PostUserSegment**](docs/UserSegmentsApi.md#postusersegment) | **Post** /segments/{projectKey}/{environmentKey} | Creates a new user segment.
*UserSegmentsApi* | [**UpdatedUnboundedSegmentTargets**](docs/UserSegmentsApi.md#updatedunboundedsegmenttargets) | **Post** /segments/{projectKey}/{environmentKey}/{userSegmentKey}/unbounded-users | Update targets included or excluded in an unbounded segment
*UserSettingsApi* | [**GetExpiringUserTargetsForUser**](docs/UserSettingsApi.md#getexpiringusertargetsforuser) | **Get** /users/{projectKey}/{userKey}/expiring-user-targets/{environmentKey} | Get expiring dates on flags for user
*UserSettingsApi* | [**GetUserFlagSetting**](docs/UserSettingsApi.md#getuserflagsetting) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Fetch a single flag setting for a user by key.
*UserSettingsApi* | [**GetUserFlagSettings**](docs/UserSettingsApi.md#getuserflagsettings) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags | Fetch a single flag setting for a user by key.
*UserSettingsApi* | [**PatchExpiringUserTargetsForFlags**](docs/UserSettingsApi.md#patchexpiringusertargetsforflags) | **Patch** /users/{projectKey}/{userKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets for a single user on all flags
*UserSettingsApi* | [**PutFlagSetting**](docs/UserSettingsApi.md#putflagsetting) | **Put** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Specifically enable or disable a feature flag for a user based on their key.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{projectKey}/{environmentKey}/{userKey} | Delete a user by ID.
*UsersApi* | [**GetSearchUsers**](docs/UsersApi.md#getsearchusers) | **Get** /user-search/{projectKey}/{environmentKey} | Search users in LaunchDarkly based on their last active date, or a search query. It should not be used to enumerate all users in LaunchDarkly-- use the List users API resource.
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{projectKey}/{environmentKey}/{userKey} | Get a user by key.
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /users/{projectKey}/{environmentKey} | List all users in the environment. Includes the total count of users. In each page, there will be up to &#39;limit&#39; users returned (default 20). This is useful for exporting all users in the system for further analysis. Paginated collections will include a next link containing a URL with the next set of elements in the collection.
*WebhooksApi* | [**DeleteWebhook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhooks/{resourceId} | Delete a webhook by ID.
*WebhooksApi* | [**GetWebhook**](docs/WebhooksApi.md#getwebhook) | **Get** /webhooks/{resourceId} | Get a webhook by ID.
*WebhooksApi* | [**GetWebhooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhooks | Fetch a list of all webhooks.
*WebhooksApi* | [**PatchWebhook**](docs/WebhooksApi.md#patchwebhook) | **Patch** /webhooks/{resourceId} | Modify a webhook by ID.
*WebhooksApi* | [**PostWebhook**](docs/WebhooksApi.md#postwebhook) | **Post** /webhooks | Create a webhook.


## Documentation For Models

 - [ApprovalRequest](docs/ApprovalRequest.md)
 - [ApprovalRequestApplyConfigBody](docs/ApprovalRequestApplyConfigBody.md)
 - [ApprovalRequestConfigBody](docs/ApprovalRequestConfigBody.md)
 - [ApprovalRequestReview](docs/ApprovalRequestReview.md)
 - [ApprovalRequestReviewConfigBody](docs/ApprovalRequestReviewConfigBody.md)
 - [ApprovalRequestReviewStatus](docs/ApprovalRequestReviewStatus.md)
 - [ApprovalRequests](docs/ApprovalRequests.md)
 - [AuditLogEntries](docs/AuditLogEntries.md)
 - [AuditLogEntry](docs/AuditLogEntry.md)
 - [AuditLogEntryTarget](docs/AuditLogEntryTarget.md)
 - [Clause](docs/Clause.md)
 - [ClientSideAvailability](docs/ClientSideAvailability.md)
 - [CopyActions](docs/CopyActions.md)
 - [CustomProperty](docs/CustomProperty.md)
 - [CustomPropertyValues](docs/CustomPropertyValues.md)
 - [CustomRole](docs/CustomRole.md)
 - [CustomRoleBody](docs/CustomRoleBody.md)
 - [CustomRoles](docs/CustomRoles.md)
 - [Defaults](docs/Defaults.md)
 - [Destination](docs/Destination.md)
 - [DestinationAmazonKinesis](docs/DestinationAmazonKinesis.md)
 - [DestinationBody](docs/DestinationBody.md)
 - [DestinationGooglePubSub](docs/DestinationGooglePubSub.md)
 - [DestinationMParticle](docs/DestinationMParticle.md)
 - [DestinationSegment](docs/DestinationSegment.md)
 - [Destinations](docs/Destinations.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentApprovalSettings](docs/EnvironmentApprovalSettings.md)
 - [EnvironmentPost](docs/EnvironmentPost.md)
 - [EvaluationUsageError](docs/EvaluationUsageError.md)
 - [Events](docs/Events.md)
 - [FeatureFlag](docs/FeatureFlag.md)
 - [FeatureFlagBody](docs/FeatureFlagBody.md)
 - [FeatureFlagConfig](docs/FeatureFlagConfig.md)
 - [FeatureFlagCopyBody](docs/FeatureFlagCopyBody.md)
 - [FeatureFlagCopyObject](docs/FeatureFlagCopyObject.md)
 - [FeatureFlagScheduledChange](docs/FeatureFlagScheduledChange.md)
 - [FeatureFlagScheduledChanges](docs/FeatureFlagScheduledChanges.md)
 - [FeatureFlagScheduledChangesConflicts](docs/FeatureFlagScheduledChangesConflicts.md)
 - [FeatureFlagScheduledChangesConflictsInstructions](docs/FeatureFlagScheduledChangesConflictsInstructions.md)
 - [FeatureFlagStatus](docs/FeatureFlagStatus.md)
 - [FeatureFlagStatusAcrossEnvironments](docs/FeatureFlagStatusAcrossEnvironments.md)
 - [FeatureFlagStatusForQueriedEnvironment](docs/FeatureFlagStatusForQueriedEnvironment.md)
 - [FeatureFlagStatuses](docs/FeatureFlagStatuses.md)
 - [FeatureFlags](docs/FeatureFlags.md)
 - [FlagConfigScheduledChangesConflictsBody](docs/FlagConfigScheduledChangesConflictsBody.md)
 - [FlagConfigScheduledChangesPatchBody](docs/FlagConfigScheduledChangesPatchBody.md)
 - [FlagConfigScheduledChangesPostBody](docs/FlagConfigScheduledChangesPostBody.md)
 - [FlagListItem](docs/FlagListItem.md)
 - [HierarchicalLinks](docs/HierarchicalLinks.md)
 - [Id](docs/Id.md)
 - [Integration](docs/Integration.md)
 - [IntegrationLinks](docs/IntegrationLinks.md)
 - [IntegrationSubscription](docs/IntegrationSubscription.md)
 - [IntegrationSubscriptionStatus](docs/IntegrationSubscriptionStatus.md)
 - [Integrations](docs/Integrations.md)
 - [Link](docs/Link.md)
 - [Links](docs/Links.md)
 - [MaUbyCategory](docs/MaUbyCategory.md)
 - [Mau](docs/Mau.md)
 - [MauMetadata](docs/MauMetadata.md)
 - [Member](docs/Member.md)
 - [MemberLastSeenMetadata](docs/MemberLastSeenMetadata.md)
 - [Members](docs/Members.md)
 - [MembersBody](docs/MembersBody.md)
 - [ModelFallthrough](docs/ModelFallthrough.md)
 - [PatchComment](docs/PatchComment.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [Policy](docs/Policy.md)
 - [Prerequisite](docs/Prerequisite.md)
 - [Project](docs/Project.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [Projects](docs/Projects.md)
 - [RelayProxyConfig](docs/RelayProxyConfig.md)
 - [RelayProxyConfigBody](docs/RelayProxyConfigBody.md)
 - [RelayProxyConfigs](docs/RelayProxyConfigs.md)
 - [Role](docs/Role.md)
 - [Rollout](docs/Rollout.md)
 - [Rule](docs/Rule.md)
 - [ScheduledChangesFeatureFlagConflict](docs/ScheduledChangesFeatureFlagConflict.md)
 - [SemanticPatchInstruction](docs/SemanticPatchInstruction.md)
 - [SemanticPatchInstructionInner](docs/SemanticPatchInstructionInner.md)
 - [SemanticPatchOperation](docs/SemanticPatchOperation.md)
 - [Site](docs/Site.md)
 - [Statement](docs/Statement.md)
 - [Stream](docs/Stream.md)
 - [StreamBySdk](docs/StreamBySdk.md)
 - [StreamBySdkLinks](docs/StreamBySdkLinks.md)
 - [StreamBySdkLinksMetadata](docs/StreamBySdkLinksMetadata.md)
 - [StreamLinks](docs/StreamLinks.md)
 - [StreamSdkVersion](docs/StreamSdkVersion.md)
 - [StreamSdkVersionData](docs/StreamSdkVersionData.md)
 - [StreamUsageError](docs/StreamUsageError.md)
 - [StreamUsageLinks](docs/StreamUsageLinks.md)
 - [StreamUsageMetadata](docs/StreamUsageMetadata.md)
 - [StreamUsageSeries](docs/StreamUsageSeries.md)
 - [Streams](docs/Streams.md)
 - [SubscriptionBody](docs/SubscriptionBody.md)
 - [Target](docs/Target.md)
 - [Token](docs/Token.md)
 - [TokenBody](docs/TokenBody.md)
 - [Tokens](docs/Tokens.md)
 - [UnboundedSegmentTargetChanges](docs/UnboundedSegmentTargetChanges.md)
 - [UnboundedSegmentTargetsBody](docs/UnboundedSegmentTargetsBody.md)
 - [Usage](docs/Usage.md)
 - [UsageError](docs/UsageError.md)
 - [UsageLinks](docs/UsageLinks.md)
 - [User](docs/User.md)
 - [UserFlagSetting](docs/UserFlagSetting.md)
 - [UserFlagSettings](docs/UserFlagSettings.md)
 - [UserRecord](docs/UserRecord.md)
 - [UserSegment](docs/UserSegment.md)
 - [UserSegmentBody](docs/UserSegmentBody.md)
 - [UserSegmentRule](docs/UserSegmentRule.md)
 - [UserSegments](docs/UserSegments.md)
 - [UserSettingsBody](docs/UserSettingsBody.md)
 - [UserTargetingExpirationForFlag](docs/UserTargetingExpirationForFlag.md)
 - [UserTargetingExpirationForFlags](docs/UserTargetingExpirationForFlags.md)
 - [UserTargetingExpirationForSegment](docs/UserTargetingExpirationForSegment.md)
 - [UserTargetingExpirationOnFlagsForUser](docs/UserTargetingExpirationOnFlagsForUser.md)
 - [UserTargetingExpirationResourceIdForFlag](docs/UserTargetingExpirationResourceIdForFlag.md)
 - [Users](docs/Users.md)
 - [Variation](docs/Variation.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookBody](docs/WebhookBody.md)
 - [Webhooks](docs/Webhooks.md)
 - [WeightedVariation](docs/WeightedVariation.md)


## Author

support@launchdarkly.com

## Sample Code

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/launchdarkly/api-client-go"
)

func main() {
	apiKey := os.Getenv("LD_API_KEY")
	if apiKey == "" {
		panic("LD_API_KEY env var was empty!")
	}
	client := ldapi.NewAPIClient(ldapi.NewConfiguration())
	ctx := context.WithValue(context.Background(), ldapi.ContextAPIKey, ldapi.APIKey{
		Key: apiKey,
	})

	// Create a multi-variate feature flag
	body := ldapi.FeatureFlagBody{
		Name: "Test Flag Go",
		Key:  "test-go",
		Variations: []ldapi.Variation{
			{Value: []int{1, 2}},
			{Value: []int{3, 4}},
			{Value: []int{5}}
	flag, _, err := client.FeatureFlagsApi.PostFeatureFlag(ctx, "openapi", body, nil)
	if err != nil {
		panic(fmt.Errorf("create failed: %s", err))
	}
	fmt.Printf("Created flag: %+v\n", flag)
	// Clean up new flag
	defer func() {
		if _, err := client.FeatureFlagsApi.DeleteFeatureFlag(ctx, "openapi", body.Key); err != nil {
			panic(fmt.Errorf("delete failed: %s", err))
		}
	}()
}
```
