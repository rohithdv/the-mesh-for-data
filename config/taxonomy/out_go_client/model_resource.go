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

// Resource Data set or other type of resource.
type Resource struct {
	// Name of the data set
	Name string `json:"name"`
	// location of dataset credentials
	Creds *string `json:"creds,omitempty"`
	Tags *map[string]map[string]interface{} `json:"tags,omitempty"`
	// List of column names in the data set with their associated tags. They must be key value pairs.
	Columns *[]ResourceColumns `json:"columns,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(name string) *Resource {
	this := Resource{}
	this.Name = name
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetName returns the Name field value
func (o *Resource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Resource) SetName(v string) {
	o.Name = v
}

// GetCreds returns the Creds field value if set, zero value otherwise.
func (o *Resource) GetCreds() string {
	if o == nil || o.Creds == nil {
		var ret string
		return ret
	}
	return *o.Creds
}

// GetCredsOk returns a tuple with the Creds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetCredsOk() (*string, bool) {
	if o == nil || o.Creds == nil {
		return nil, false
	}
	return o.Creds, true
}

// HasCreds returns a boolean if a field has been set.
func (o *Resource) HasCreds() bool {
	if o != nil && o.Creds != nil {
		return true
	}

	return false
}

// SetCreds gets a reference to the given string and assigns it to the Creds field.
func (o *Resource) SetCreds(v string) {
	o.Creds = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Resource) GetTags() map[string]map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetTagsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Resource) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]map[string]interface{} and assigns it to the Tags field.
func (o *Resource) SetTags(v map[string]map[string]interface{}) {
	o.Tags = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *Resource) GetColumns() []ResourceColumns {
	if o == nil || o.Columns == nil {
		var ret []ResourceColumns
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetColumnsOk() (*[]ResourceColumns, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *Resource) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ResourceColumns and assigns it to the Columns field.
func (o *Resource) SetColumns(v []ResourceColumns) {
	o.Columns = &v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Creds != nil {
		toSerialize["creds"] = o.Creds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

