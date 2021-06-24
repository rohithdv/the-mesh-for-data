/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v2opa

import (
	"net/http"
	"strings"

	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"GetPoliciesDecisions",
			strings.ToUpper("Get"),
			"/getPoliciesDecisions",
			c.GetPoliciesDecisions,
		},
	}
}

// GetPoliciesDecisions - getPoliciesDecisions
func (c *DefaultApiController) GetPoliciesDecisions(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	// input := strings.Split(query.Get("input"), ",")

	input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("Name_example"))} // []PolicymanagerRequest | input values that need to be considered for filter

	result, err := c.service.GetPoliciesDecisions(r.Context(), input)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
