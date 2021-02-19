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

type UserTargetingExpirationForFlag struct {
	// Unix epoch time in milliseconds specifying the expiration date
	ExpirationDate int64 `json:"expirationDate,omitempty"`
	// the ID of the variation that the user is targeted on a flag
	VariationId string `json:"variationId,omitempty"`
	// Unique identifier for the user
	UserKey string `json:"userKey,omitempty"`
	Id string `json:"_id,omitempty"`
	ResourceId *UserTargetingExpirationResourceIdForFlag `json:"_resourceId,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Version int32 `json:"_version,omitempty"`
}
