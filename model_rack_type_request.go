/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2.2 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the RackTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackTypeRequest{}

// RackTypeRequest Adds support for custom fields and tags.
type RackTypeRequest struct {
	Manufacturer BriefManufacturerRequest                     `json:"manufacturer"`
	Model        string                                       `json:"model"`
	Slug         string                                       `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description  *string                                      `json:"description,omitempty"`
	FormFactor   NullablePatchedWritableRackRequestFormFactor `json:"form_factor,omitempty"`
	Width        *RackWidthValue                              `json:"width,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Starting unit for rack
	StartingUnit *int32 `json:"starting_unit,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32                               `json:"outer_depth,omitempty"`
	OuterUnit  NullablePatchedWritableRackRequestOuterUnit `json:"outer_unit,omitempty"`
	Weight     NullableFloat64                             `json:"weight,omitempty"`
	// Maximum load capacity for the rack
	MaxWeight  NullableInt32                       `json:"max_weight,omitempty"`
	WeightUnit NullableDeviceTypeRequestWeightUnit `json:"weight_unit,omitempty"`
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth        NullableInt32          `json:"mounting_depth,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackTypeRequest RackTypeRequest

// NewRackTypeRequest instantiates a new RackTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackTypeRequest(manufacturer BriefManufacturerRequest, model string, slug string) *RackTypeRequest {
	this := RackTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	return &this
}

// NewRackTypeRequestWithDefaults instantiates a new RackTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackTypeRequestWithDefaults() *RackTypeRequest {
	this := RackTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *RackTypeRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil {
		var ret BriefManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *RackTypeRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *RackTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RackTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *RackTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *RackTypeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RackTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RackTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RackTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetFormFactor() PatchedWritableRackRequestFormFactor {
	if o == nil || IsNil(o.FormFactor.Get()) {
		var ret PatchedWritableRackRequestFormFactor
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetFormFactorOk() (*PatchedWritableRackRequestFormFactor, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *RackTypeRequest) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullablePatchedWritableRackRequestFormFactor and assigns it to the FormFactor field.
func (o *RackTypeRequest) SetFormFactor(v PatchedWritableRackRequestFormFactor) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *RackTypeRequest) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *RackTypeRequest) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *RackTypeRequest) GetWidth() RackWidthValue {
	if o == nil || IsNil(o.Width) {
		var ret RackWidthValue
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetWidthOk() (*RackWidthValue, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *RackTypeRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given RackWidthValue and assigns it to the Width field.
func (o *RackTypeRequest) SetWidth(v RackWidthValue) {
	o.Width = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *RackTypeRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *RackTypeRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *RackTypeRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetStartingUnit returns the StartingUnit field value if set, zero value otherwise.
func (o *RackTypeRequest) GetStartingUnit() int32 {
	if o == nil || IsNil(o.StartingUnit) {
		var ret int32
		return ret
	}
	return *o.StartingUnit
}

// GetStartingUnitOk returns a tuple with the StartingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetStartingUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingUnit) {
		return nil, false
	}
	return o.StartingUnit, true
}

// HasStartingUnit returns a boolean if a field has been set.
func (o *RackTypeRequest) HasStartingUnit() bool {
	if o != nil && !IsNil(o.StartingUnit) {
		return true
	}

	return false
}

// SetStartingUnit gets a reference to the given int32 and assigns it to the StartingUnit field.
func (o *RackTypeRequest) SetStartingUnit(v int32) {
	o.StartingUnit = &v
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *RackTypeRequest) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *RackTypeRequest) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *RackTypeRequest) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *RackTypeRequest) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *RackTypeRequest) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}

// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *RackTypeRequest) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *RackTypeRequest) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *RackTypeRequest) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *RackTypeRequest) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}

// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *RackTypeRequest) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *RackTypeRequest) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetOuterUnit() PatchedWritableRackRequestOuterUnit {
	if o == nil || IsNil(o.OuterUnit.Get()) {
		var ret PatchedWritableRackRequestOuterUnit
		return ret
	}
	return *o.OuterUnit.Get()
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetOuterUnitOk() (*PatchedWritableRackRequestOuterUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterUnit.Get(), o.OuterUnit.IsSet()
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *RackTypeRequest) HasOuterUnit() bool {
	if o != nil && o.OuterUnit.IsSet() {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given NullablePatchedWritableRackRequestOuterUnit and assigns it to the OuterUnit field.
func (o *RackTypeRequest) SetOuterUnit(v PatchedWritableRackRequestOuterUnit) {
	o.OuterUnit.Set(&v)
}

// SetOuterUnitNil sets the value for OuterUnit to be an explicit nil
func (o *RackTypeRequest) SetOuterUnitNil() {
	o.OuterUnit.Set(nil)
}

// UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
func (o *RackTypeRequest) UnsetOuterUnit() {
	o.OuterUnit.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *RackTypeRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *RackTypeRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *RackTypeRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *RackTypeRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetMaxWeight() int32 {
	if o == nil || IsNil(o.MaxWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWeight.Get()
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetMaxWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxWeight.Get(), o.MaxWeight.IsSet()
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *RackTypeRequest) HasMaxWeight() bool {
	if o != nil && o.MaxWeight.IsSet() {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given NullableInt32 and assigns it to the MaxWeight field.
func (o *RackTypeRequest) SetMaxWeight(v int32) {
	o.MaxWeight.Set(&v)
}

// SetMaxWeightNil sets the value for MaxWeight to be an explicit nil
func (o *RackTypeRequest) SetMaxWeightNil() {
	o.MaxWeight.Set(nil)
}

// UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
func (o *RackTypeRequest) UnsetMaxWeight() {
	o.MaxWeight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetWeightUnit() DeviceTypeRequestWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeRequestWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *RackTypeRequest) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeRequestWeightUnit and assigns it to the WeightUnit field.
func (o *RackTypeRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *RackTypeRequest) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *RackTypeRequest) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetMountingDepth returns the MountingDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackTypeRequest) GetMountingDepth() int32 {
	if o == nil || IsNil(o.MountingDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.MountingDepth.Get()
}

// GetMountingDepthOk returns a tuple with the MountingDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackTypeRequest) GetMountingDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountingDepth.Get(), o.MountingDepth.IsSet()
}

// HasMountingDepth returns a boolean if a field has been set.
func (o *RackTypeRequest) HasMountingDepth() bool {
	if o != nil && o.MountingDepth.IsSet() {
		return true
	}

	return false
}

// SetMountingDepth gets a reference to the given NullableInt32 and assigns it to the MountingDepth field.
func (o *RackTypeRequest) SetMountingDepth(v int32) {
	o.MountingDepth.Set(&v)
}

// SetMountingDepthNil sets the value for MountingDepth to be an explicit nil
func (o *RackTypeRequest) SetMountingDepthNil() {
	o.MountingDepth.Set(nil)
}

// UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
func (o *RackTypeRequest) UnsetMountingDepth() {
	o.MountingDepth.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *RackTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *RackTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *RackTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackTypeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *RackTypeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o RackTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.FormFactor.IsSet() {
		toSerialize["form_factor"] = o.FormFactor.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.StartingUnit) {
		toSerialize["starting_unit"] = o.StartingUnit
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if o.OuterUnit.IsSet() {
		toSerialize["outer_unit"] = o.OuterUnit.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.MaxWeight.IsSet() {
		toSerialize["max_weight"] = o.MaxWeight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if o.MountingDepth.IsSet() {
		toSerialize["mounting_depth"] = o.MountingDepth.Get()
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

func (o *RackTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
		"slug",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varRackTypeRequest := _RackTypeRequest{}

	err = json.Unmarshal(data, &varRackTypeRequest)

	if err != nil {
		return err
	}

	*o = RackTypeRequest(varRackTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "form_factor")
		delete(additionalProperties, "width")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "starting_unit")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "mounting_depth")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackTypeRequest struct {
	value *RackTypeRequest
	isSet bool
}

func (v NullableRackTypeRequest) Get() *RackTypeRequest {
	return v.value
}

func (v *NullableRackTypeRequest) Set(val *RackTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRackTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRackTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackTypeRequest(val *RackTypeRequest) *NullableRackTypeRequest {
	return &NullableRackTypeRequest{value: val, isSet: true}
}

func (v NullableRackTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
