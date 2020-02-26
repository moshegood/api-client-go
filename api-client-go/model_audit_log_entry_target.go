/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.30
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type AuditLogEntryTarget struct {
	Links *Links `json:"_links,omitempty"`
	Name string `json:"name,omitempty"`
	Resources []string `json:"resources,omitempty"`
}
