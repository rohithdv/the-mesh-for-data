/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PolicymanagerResponse This taxonomy defines the structure of response received from policymanager with enforcements.
type PolicymanagerResponse struct {
	// This is the id returned by the governance engine
	DecisionId *string `json:"decision_id,omitempty"`
	// While showing the result, action contains the action type and the associated entity on which action has been taken.
	Result []PolicymanagerResponseResult `json:"result"`
}

// NewPolicymanagerResponse instantiates a new PolicymanagerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicymanagerResponse(result []PolicymanagerResponseResult) *PolicymanagerResponse {
	this := PolicymanagerResponse{}
	this.Result = result
	return &this
}

// NewPolicymanagerResponseWithDefaults instantiates a new PolicymanagerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicymanagerResponseWithDefaults() *PolicymanagerResponse {
	this := PolicymanagerResponse{}
	return &this
}

// GetDecisionId returns the DecisionId field value if set, zero value otherwise.
func (o *PolicymanagerResponse) GetDecisionId() string {
	if o == nil || o.DecisionId == nil {
		var ret string
		return ret
	}
	return *o.DecisionId
}

// GetDecisionIdOk returns a tuple with the DecisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicymanagerResponse) GetDecisionIdOk() (*string, bool) {
	if o == nil || o.DecisionId == nil {
		return nil, false
	}
	return o.DecisionId, true
}

// HasDecisionId returns a boolean if a field has been set.
func (o *PolicymanagerResponse) HasDecisionId() bool {
	if o != nil && o.DecisionId != nil {
		return true
	}

	return false
}

// SetDecisionId gets a reference to the given string and assigns it to the DecisionId field.
func (o *PolicymanagerResponse) SetDecisionId(v string) {
	o.DecisionId = &v
}

// GetResult returns the Result field value
func (o *PolicymanagerResponse) GetResult() []PolicymanagerResponseResult {
	if o == nil {
		var ret []PolicymanagerResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PolicymanagerResponse) GetResultOk() (*[]PolicymanagerResponseResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PolicymanagerResponse) SetResult(v []PolicymanagerResponseResult) {
	o.Result = v
}

func (o PolicymanagerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DecisionId != nil {
		toSerialize["decision_id"] = o.DecisionId
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullablePolicymanagerResponse struct {
	value *PolicymanagerResponse
	isSet bool
}

func (v NullablePolicymanagerResponse) Get() *PolicymanagerResponse {
	return v.value
}

func (v *NullablePolicymanagerResponse) Set(val *PolicymanagerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicymanagerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicymanagerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicymanagerResponse(val *PolicymanagerResponse) *NullablePolicymanagerResponse {
	return &NullablePolicymanagerResponse{value: val, isSet: true}
}

func (v NullablePolicymanagerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicymanagerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


