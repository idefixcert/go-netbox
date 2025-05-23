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
	"time"
)

// checks if the VRF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VRF{}

// VRF Adds support for custom fields and tags.
type VRF struct {
	Id         int32   `json:"id"`
	Url        string  `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display    string  `json:"display"`
	Name       string  `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd     NullableString      `json:"rd,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique        *bool                  `json:"enforce_unique,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	ImportTargets        []RouteTarget          `json:"import_targets,omitempty"`
	ExportTargets        []RouteTarget          `json:"export_targets,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created,omitempty"`
	LastUpdated          NullableTime           `json:"last_updated,omitempty"`
	IpaddressCount       *int64                 `json:"ipaddress_count,omitempty"`
	PrefixCount          *int64                 `json:"prefix_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VRF VRF

// NewVRF instantiates a new VRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVRF(id int32, url string, display string, name string) *VRF {
	this := VRF{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewVRFWithDefaults instantiates a new VRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVRFWithDefaults() *VRF {
	this := VRF{}
	return &this
}

// GetId returns the Id field value
func (o *VRF) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VRF) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VRF) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VRF) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VRF) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VRF) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *VRF) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *VRF) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *VRF) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *VRF) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VRF) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VRF) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *VRF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VRF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VRF) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetRd() string {
	if o == nil || IsNil(o.Rd.Get()) {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *VRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *VRF) SetRd(v string) {
	o.Rd.Set(&v)
}

// SetRdNil sets the value for Rd to be an explicit nil
func (o *VRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *VRF) UnsetRd() {
	o.Rd.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VRF) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *VRF) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VRF) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VRF) UnsetTenant() {
	o.Tenant.Unset()
}

// GetEnforceUnique returns the EnforceUnique field value if set, zero value otherwise.
func (o *VRF) GetEnforceUnique() bool {
	if o == nil || IsNil(o.EnforceUnique) {
		var ret bool
		return ret
	}
	return *o.EnforceUnique
}

// GetEnforceUniqueOk returns a tuple with the EnforceUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetEnforceUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceUnique) {
		return nil, false
	}
	return o.EnforceUnique, true
}

// HasEnforceUnique returns a boolean if a field has been set.
func (o *VRF) HasEnforceUnique() bool {
	if o != nil && !IsNil(o.EnforceUnique) {
		return true
	}

	return false
}

// SetEnforceUnique gets a reference to the given bool and assigns it to the EnforceUnique field.
func (o *VRF) SetEnforceUnique(v bool) {
	o.EnforceUnique = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VRF) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VRF) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VRF) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VRF) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VRF) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VRF) SetComments(v string) {
	o.Comments = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *VRF) GetImportTargets() []RouteTarget {
	if o == nil || IsNil(o.ImportTargets) {
		var ret []RouteTarget
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetImportTargetsOk() ([]RouteTarget, bool) {
	if o == nil || IsNil(o.ImportTargets) {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *VRF) HasImportTargets() bool {
	if o != nil && !IsNil(o.ImportTargets) {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []RouteTarget and assigns it to the ImportTargets field.
func (o *VRF) SetImportTargets(v []RouteTarget) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *VRF) GetExportTargets() []RouteTarget {
	if o == nil || IsNil(o.ExportTargets) {
		var ret []RouteTarget
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetExportTargetsOk() ([]RouteTarget, bool) {
	if o == nil || IsNil(o.ExportTargets) {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *VRF) HasExportTargets() bool {
	if o != nil && !IsNil(o.ExportTargets) {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []RouteTarget and assigns it to the ExportTargets field.
func (o *VRF) SetExportTargets(v []RouteTarget) {
	o.ExportTargets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VRF) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VRF) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VRF) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VRF) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VRF) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VRF) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *VRF) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *VRF) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *VRF) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *VRF) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VRF) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *VRF) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *VRF) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *VRF) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetIpaddressCount returns the IpaddressCount field value if set, zero value otherwise.
func (o *VRF) GetIpaddressCount() int64 {
	if o == nil || IsNil(o.IpaddressCount) {
		var ret int64
		return ret
	}
	return *o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetIpaddressCountOk() (*int64, bool) {
	if o == nil || IsNil(o.IpaddressCount) {
		return nil, false
	}
	return o.IpaddressCount, true
}

// HasIpaddressCount returns a boolean if a field has been set.
func (o *VRF) HasIpaddressCount() bool {
	if o != nil && !IsNil(o.IpaddressCount) {
		return true
	}

	return false
}

// SetIpaddressCount gets a reference to the given int64 and assigns it to the IpaddressCount field.
func (o *VRF) SetIpaddressCount(v int64) {
	o.IpaddressCount = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *VRF) GetPrefixCount() int64 {
	if o == nil || IsNil(o.PrefixCount) {
		var ret int64
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetPrefixCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixCount) {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *VRF) HasPrefixCount() bool {
	if o != nil && !IsNil(o.PrefixCount) {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int64 and assigns it to the PrefixCount field.
func (o *VRF) SetPrefixCount(v int64) {
	o.PrefixCount = &v
}

func (o VRF) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VRF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.EnforceUnique) {
		toSerialize["enforce_unique"] = o.EnforceUnique
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.ImportTargets) {
		toSerialize["import_targets"] = o.ImportTargets
	}
	if !IsNil(o.ExportTargets) {
		toSerialize["export_targets"] = o.ExportTargets
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.LastUpdated.IsSet() {
		toSerialize["last_updated"] = o.LastUpdated.Get()
	}
	if !IsNil(o.IpaddressCount) {
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if !IsNil(o.PrefixCount) {
		toSerialize["prefix_count"] = o.PrefixCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VRF) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
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
	varVRF := _VRF{}

	err = json.Unmarshal(data, &varVRF)

	if err != nil {
		return err
	}

	*o = VRF(varVRF)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rd")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "enforce_unique")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "import_targets")
		delete(additionalProperties, "export_targets")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "ipaddress_count")
		delete(additionalProperties, "prefix_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVRF struct {
	value *VRF
	isSet bool
}

func (v NullableVRF) Get() *VRF {
	return v.value
}

func (v *NullableVRF) Set(val *VRF) {
	v.value = val
	v.isSet = true
}

func (v NullableVRF) IsSet() bool {
	return v.isSet
}

func (v *NullableVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVRF(val *VRF) *NullableVRF {
	return &NullableVRF{value: val, isSet: true}
}

func (v NullableVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
