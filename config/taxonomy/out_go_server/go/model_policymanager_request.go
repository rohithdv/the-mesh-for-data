/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PolicymanagerRequest - This taxonomy defines the structure sent to Open Policy Agent policy engine, which determines whether or how the data requested may be used by the requester.
type PolicymanagerRequest struct {

	RequestContext RequestContext `json:"request_context,omitempty"`

	Action Action `json:"action"`

	Resource Resource `json:"resource"`
}