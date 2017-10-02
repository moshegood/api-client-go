/* 
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type FeatureFlagBody struct {

	// A human-friendly name for the feature flag. Remember to note if this flag is intended to be temporary or permanent.
	Name string `json:"name"`

	// A unique key that will be used to reference the flag in your code
	Key string `json:"key"`

	// An array of possible variations for the flag.
	Variations []Variation `json:"variations"`

	// Whether or not the flag is a temporary flag
	Temporary bool `json:"temporary,omitempty"`

	// Tags for the feature flag
	Tags []string `json:"tags,omitempty"`

	// Whether or not this flag should be made available to the client-side JavaScript SDK
	IncludeInSnippet bool `json:"includeInSnippet,omitempty"`
}