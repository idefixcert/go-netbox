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

// checks if the PatchedRackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRackReservationRequest{}

// PatchedRackReservationRequest Adds support for custom fields and tags.
type PatchedRackReservationRequest struct {
	Rack                 *BriefRackRequest          `json:"rack,omitempty"`
	Units                []int32                    `json:"units,omitempty"`
	User                 *BriefUserRequest          `json:"user,omitempty"`
	Tenant               NullableBriefTenantRequest `json:"tenant,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Comments             *string                    `json:"comments,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedRackReservationRequest PatchedRackReservationRequest

// NewPatchedRackReservationRequest instantiates a new PatchedRackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRackReservationRequest() *PatchedRackReservationRequest {
	this := PatchedRackReservationRequest{}
	return &this
}

// NewPatchedRackReservationRequestWithDefaults instantiates a new PatchedRackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRackReservationRequestWithDefaults() *PatchedRackReservationRequest {
	this := PatchedRackReservationRequest{}
	return &this
}

// GetRack returns the Rack field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetRack() BriefRackRequest {
	if o == nil || IsNil(o.Rack) {
		var ret BriefRackRequest
		return ret
	}
	return *o.Rack
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetRackOk() (*BriefRackRequest, bool) {
	if o == nil || IsNil(o.Rack) {
		return nil, false
	}
	return o.Rack, true
}

// HasRack returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasRack() bool {
	if o != nil && !IsNil(o.Rack) {
		return true
	}

	return false
}

// SetRack gets a reference to the given BriefRackRequest and assigns it to the Rack field.
func (o *PatchedRackReservationRequest) SetRack(v BriefRackRequest) {
	o.Rack = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetUnits() []int32 {
	if o == nil || IsNil(o.Units) {
		var ret []int32
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetUnitsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []int32 and assigns it to the Units field.
func (o *PatchedRackReservationRequest) SetUnits(v []int32) {
	o.Units = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetUser() BriefUserRequest {
	if o == nil || IsNil(o.User) {
		var ret BriefUserRequest
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetUserOk() (*BriefUserRequest, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given BriefUserRequest and assigns it to the User field.
func (o *PatchedRackReservationRequest) SetUser(v BriefUserRequest) {
	o.User = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRackReservationRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRackReservationRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *PatchedRackReservationRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedRackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedRackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedRackReservationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedRackReservationRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedRackReservationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedRackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedRackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedRackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedRackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rack) {
		toSerialize["rack"] = o.Rack
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *PatchedRackReservationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedRackReservationRequest := _PatchedRackReservationRequest{}

	err = json.Unmarshal(data, &varPatchedRackReservationRequest)

	if err != nil {
		return err
	}

	*o = PatchedRackReservationRequest(varPatchedRackReservationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rack")
		delete(additionalProperties, "units")
		delete(additionalProperties, "user")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedRackReservationRequest struct {
	value *PatchedRackReservationRequest
	isSet bool
}

func (v NullablePatchedRackReservationRequest) Get() *PatchedRackReservationRequest {
	return v.value
}

func (v *NullablePatchedRackReservationRequest) Set(val *PatchedRackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRackReservationRequest(val *PatchedRackReservationRequest) *NullablePatchedRackReservationRequest {
	return &NullablePatchedRackReservationRequest{value: val, isSet: true}
}

func (v NullablePatchedRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
