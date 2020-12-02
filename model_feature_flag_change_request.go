/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.9.1
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagChangeRequest struct {
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Version int32 `json:"_version,omitempty"`
	// A unix epoch time in milliseconds specifying the date the change request was requested
	CreationDate int32 `json:"creationDate,omitempty"`
	// The id of the member that requested the change
	RequestorId string `json:"requestorId,omitempty"`
	ReviewStatus *FeatureFlagChangeRequestReviewStatus `json:"reviewStatus,omitempty"`
	// | Name     | Description | | --------:| ----------- | | pending  | the feature flag change request has not been applied yet | | completed| the feature flag change request has been applied successfully | | failed   | the feature flag change request has been applied but the changes were not applied successfully | 
	Status string `json:"status,omitempty"`
	// The id of the member that applied the change request
	AppliedByMemberID string `json:"appliedByMemberID,omitempty"`
	// A unix epoch time in milliseconds specifying the date the change request was applied
	AppliedDate int32 `json:"appliedDate,omitempty"`
	CurrentReviewsByMemberId *FeatureFlagChangeRequestReview `json:"currentReviewsByMemberId,omitempty"`
	AllReviews []FeatureFlagChangeRequestReview `json:"allReviews,omitempty"`
	NotifyMemberIds []string `json:"notifyMemberIds,omitempty"`
	Instructions *SemanticPatchInstruction `json:"instructions,omitempty"`
}
