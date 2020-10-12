/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.7.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type ScheduledChangesFeatureFlagConflict struct {
	// Feature flag scheduled change id this change will conflict with
	Id string `json:"_id,omitempty"`
	// Feature flag scheduled change conflict reason
	Reason string `json:"reason,omitempty"`
}
