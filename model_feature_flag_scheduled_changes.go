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

type FeatureFlagScheduledChanges struct {
	Links *Links `json:"_links,omitempty"`
	Items []FeatureFlagScheduledChange `json:"items,omitempty"`
}
