/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Action1 - Contains one of two action types 
type Action1 struct {

	Name AllowableActionDataset `json:"name,omitempty"`

	// Represents the column names on which the actions have been taken
	Columns []string `json:"columns"`
}
