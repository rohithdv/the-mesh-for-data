/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ActionOnColumns struct {

	Name AllowableActionColumns `json:"name,omitempty"`

	// Represents the column names on which the actions have been taken
	Columns []string `json:"columns"`
}