/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 4.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagApprovalRequest struct {
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Version int32 `json:"_version,omitempty"`
	// A unix epoch time in milliseconds specifying the date the approval request was requested
	CreationDate int32 `json:"creationDate,omitempty"`
	// The id of the member that requested the change
	RequestorId string `json:"requestorId,omitempty"`
	ReviewStatus *FeatureFlagApprovalRequestReviewStatus `json:"reviewStatus,omitempty"`
	// | Name     | Description | | --------:| ----------- | | pending  | the feature flag approval request has not been applied yet | | completed| the feature flag approval request has been applied successfully | | failed   | the feature flag approval request has been applied but the changes were not applied successfully | 
	Status string `json:"status,omitempty"`
	// The id of the member that applied the approval request
	AppliedByMemberID string `json:"appliedByMemberID,omitempty"`
	// A unix epoch time in milliseconds specifying the date the approval request was applied
	AppliedDate int32 `json:"appliedDate,omitempty"`
	AllReviews []FeatureFlagApprovalRequestReview `json:"allReviews,omitempty"`
	NotifyMemberIds []string `json:"notifyMemberIds,omitempty"`
	Instructions *SemanticPatchInstruction `json:"instructions,omitempty"`
}
