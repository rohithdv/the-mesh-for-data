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

// PolicymanagerRequest This taxonomy defines the structure sent to Open Policy Agent policy engine, which determines whether or how the data requested may be used by the requester.
type PolicymanagerRequest struct {
	RequestContext *RequestContext `json:"request_context,omitempty"`
	Action Action `json:"action"`
	Resource Resource `json:"resource"`
}

// NewPolicymanagerRequest instantiates a new PolicymanagerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicymanagerRequest(action Action, resource Resource) *PolicymanagerRequest {
	this := PolicymanagerRequest{}
	this.Action = action
	this.Resource = resource
	return &this
}

// NewPolicymanagerRequestWithDefaults instantiates a new PolicymanagerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicymanagerRequestWithDefaults() *PolicymanagerRequest {
	this := PolicymanagerRequest{}
	return &this
}

// GetRequestContext returns the RequestContext field value if set, zero value otherwise.
func (o *PolicymanagerRequest) GetRequestContext() RequestContext {
	if o == nil || o.RequestContext == nil {
		var ret RequestContext
		return ret
	}
	return *o.RequestContext
}

// GetRequestContextOk returns a tuple with the RequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicymanagerRequest) GetRequestContextOk() (*RequestContext, bool) {
	if o == nil || o.RequestContext == nil {
		return nil, false
	}
	return o.RequestContext, true
}

// HasRequestContext returns a boolean if a field has been set.
func (o *PolicymanagerRequest) HasRequestContext() bool {
	if o != nil && o.RequestContext != nil {
		return true
	}

	return false
}

// SetRequestContext gets a reference to the given RequestContext and assigns it to the RequestContext field.
func (o *PolicymanagerRequest) SetRequestContext(v RequestContext) {
	o.RequestContext = &v
}

// GetAction returns the Action field value
func (o *PolicymanagerRequest) GetAction() Action {
	if o == nil {
		var ret Action
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PolicymanagerRequest) GetActionOk() (*Action, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PolicymanagerRequest) SetAction(v Action) {
	o.Action = v
}

// GetResource returns the Resource field value
func (o *PolicymanagerRequest) GetResource() Resource {
	if o == nil {
		var ret Resource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *PolicymanagerRequest) GetResourceOk() (*Resource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *PolicymanagerRequest) SetResource(v Resource) {
	o.Resource = v
}

func (o PolicymanagerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestContext != nil {
		toSerialize["request_context"] = o.RequestContext
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullablePolicymanagerRequest struct {
	value *PolicymanagerRequest
	isSet bool
}

func (v NullablePolicymanagerRequest) Get() *PolicymanagerRequest {
	return v.value
}

func (v *NullablePolicymanagerRequest) Set(val *PolicymanagerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicymanagerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicymanagerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicymanagerRequest(val *PolicymanagerRequest) *NullablePolicymanagerRequest {
	return &NullablePolicymanagerRequest{value: val, isSet: true}
}

func (v NullablePolicymanagerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicymanagerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


