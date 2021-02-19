/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.0.1
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Stream struct {
	Links *StreamUsageLinks `json:"_links,omitempty"`
	Metadata []StreamUsageMetadata `json:"metadata,omitempty"`
	Series []StreamUsageSeries `json:"series,omitempty"`
}
