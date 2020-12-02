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

type RelayProxyConfigBody struct {
	// A human-friendly name for the relay proxy configuration
	Name string `json:"name,omitempty"`
	Policy []Policy `json:"policy,omitempty"`
}
