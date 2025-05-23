/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2.2 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedMACAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedMACAddressRequest{}

// PatchedMACAddressRequest Adds support for custom fields and tags.
type PatchedMACAddressRequest struct {
	MacAddress           *string                `json:"mac_address,omitempty"`
	AssignedObjectType   NullableString         `json:"assigned_object_type,omitempty"`
	AssignedObjectId     NullableInt64          `json:"assigned_object_id,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedMACAddressRequest PatchedMACAddressRequest

// NewPatchedMACAddressRequest instantiates a new PatchedMACAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedMACAddressRequest() *PatchedMACAddressRequest {
	this := PatchedMACAddressRequest{}
	return &this
}

// NewPatchedMACAddressRequestWithDefaults instantiates a new PatchedMACAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedMACAddressRequestWithDefaults() *PatchedMACAddressRequest {
	this := PatchedMACAddressRequest{}
	return &this
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *PatchedMACAddressRequest) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMACAddressRequest) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *PatchedMACAddressRequest) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedMACAddressRequest) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType.Get()
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedMACAddressRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectType.Get(), o.AssignedObjectType.IsSet()
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasAssignedObjectType() bool {
	if o != nil && o.AssignedObjectType.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given NullableString and assigns it to the AssignedObjectType field.
func (o *PatchedMACAddressRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType.Set(&v)
}

// SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil
func (o *PatchedMACAddressRequest) SetAssignedObjectTypeNil() {
	o.AssignedObjectType.Set(nil)
}

// UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
func (o *PatchedMACAddressRequest) UnsetAssignedObjectType() {
	o.AssignedObjectType.Unset()
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedMACAddressRequest) GetAssignedObjectId() int64 {
	if o == nil || IsNil(o.AssignedObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.AssignedObjectId.Get()
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedMACAddressRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectId.Get(), o.AssignedObjectId.IsSet()
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasAssignedObjectId() bool {
	if o != nil && o.AssignedObjectId.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given NullableInt64 and assigns it to the AssignedObjectId field.
func (o *PatchedMACAddressRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId.Set(&v)
}

// SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil
func (o *PatchedMACAddressRequest) SetAssignedObjectIdNil() {
	o.AssignedObjectId.Set(nil)
}

// UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
func (o *PatchedMACAddressRequest) UnsetAssignedObjectId() {
	o.AssignedObjectId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedMACAddressRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMACAddressRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedMACAddressRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedMACAddressRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMACAddressRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedMACAddressRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedMACAddressRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMACAddressRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedMACAddressRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedMACAddressRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMACAddressRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedMACAddressRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedMACAddressRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedMACAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedMACAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MacAddress) {
		toSerialize["mac_address"] = o.MacAddress
	}
	if o.AssignedObjectType.IsSet() {
		toSerialize["assigned_object_type"] = o.AssignedObjectType.Get()
	}
	if o.AssignedObjectId.IsSet() {
		toSerialize["assigned_object_id"] = o.AssignedObjectId.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedMACAddressRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedMACAddressRequest := _PatchedMACAddressRequest{}

	err = json.Unmarshal(data, &varPatchedMACAddressRequest)

	if err != nil {
		return err
	}

	*o = PatchedMACAddressRequest(varPatchedMACAddressRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac_address")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedMACAddressRequest struct {
	value *PatchedMACAddressRequest
	isSet bool
}

func (v NullablePatchedMACAddressRequest) Get() *PatchedMACAddressRequest {
	return v.value
}

func (v *NullablePatchedMACAddressRequest) Set(val *PatchedMACAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedMACAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedMACAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedMACAddressRequest(val *PatchedMACAddressRequest) *NullablePatchedMACAddressRequest {
	return &NullablePatchedMACAddressRequest{value: val, isSet: true}
}

func (v NullablePatchedMACAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedMACAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
