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

// Default values to be used when a new environment is created.
type Defaults struct {
	// The index of the variation to be served when a flag's targeting is on (default variation).
	OnVariation int32 `json:"onVariation"`
	// The index of the variation to be served when a flag is off.
	OffVariation int32 `json:"offVariation"`
}
