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

// checks if the PatchedWritableVMInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableVMInterfaceRequest{}

// PatchedWritableVMInterfaceRequest Adds support for custom fields and tags.
type PatchedWritableVMInterfaceRequest struct {
	VirtualMachine        *BriefVirtualMachineRequest                 `json:"virtual_machine,omitempty"`
	Name                  *string                                     `json:"name,omitempty"`
	Enabled               *bool                                       `json:"enabled,omitempty"`
	Parent                NullableInt32                               `json:"parent,omitempty"`
	Bridge                NullableInt32                               `json:"bridge,omitempty"`
	Mtu                   NullableInt32                               `json:"mtu,omitempty"`
	PrimaryMacAddress     NullableBriefMACAddressRequest              `json:"primary_mac_address,omitempty"`
	Description           *string                                     `json:"description,omitempty"`
	Mode                  NullablePatchedWritableInterfaceRequestMode `json:"mode,omitempty"`
	UntaggedVlan          NullableBriefVLANRequest                    `json:"untagged_vlan,omitempty"`
	TaggedVlans           []int32                                     `json:"tagged_vlans,omitempty"`
	QinqSvlan             NullableBriefVLANRequest                    `json:"qinq_svlan,omitempty"`
	VlanTranslationPolicy NullableBriefVLANTranslationPolicyRequest   `json:"vlan_translation_policy,omitempty"`
	Vrf                   NullableBriefVRFRequest                     `json:"vrf,omitempty"`
	Tags                  []NestedTagRequest                          `json:"tags,omitempty"`
	CustomFields          map[string]interface{}                      `json:"custom_fields,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _PatchedWritableVMInterfaceRequest PatchedWritableVMInterfaceRequest

// NewPatchedWritableVMInterfaceRequest instantiates a new PatchedWritableVMInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableVMInterfaceRequest() *PatchedWritableVMInterfaceRequest {
	this := PatchedWritableVMInterfaceRequest{}
	return &this
}

// NewPatchedWritableVMInterfaceRequestWithDefaults instantiates a new PatchedWritableVMInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableVMInterfaceRequestWithDefaults() *PatchedWritableVMInterfaceRequest {
	this := PatchedWritableVMInterfaceRequest{}
	return &this
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetVirtualMachine() BriefVirtualMachineRequest {
	if o == nil || IsNil(o.VirtualMachine) {
		var ret BriefVirtualMachineRequest
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool) {
	if o == nil || IsNil(o.VirtualMachine) {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasVirtualMachine() bool {
	if o != nil && !IsNil(o.VirtualMachine) {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given BriefVirtualMachineRequest and assigns it to the VirtualMachine field.
func (o *PatchedWritableVMInterfaceRequest) SetVirtualMachine(v BriefVirtualMachineRequest) {
	o.VirtualMachine = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableVMInterfaceRequest) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableVMInterfaceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *PatchedWritableVMInterfaceRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetBridge() int32 {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret int32
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetBridgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableInt32 and assigns it to the Bridge field.
func (o *PatchedWritableVMInterfaceRequest) SetBridge(v int32) {
	o.Bridge.Set(&v)
}

// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *PatchedWritableVMInterfaceRequest) SetMtu(v int32) {
	o.Mtu.Set(&v)
}

// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetMtu() {
	o.Mtu.Unset()
}

// GetPrimaryMacAddress returns the PrimaryMacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetPrimaryMacAddress() BriefMACAddressRequest {
	if o == nil || IsNil(o.PrimaryMacAddress.Get()) {
		var ret BriefMACAddressRequest
		return ret
	}
	return *o.PrimaryMacAddress.Get()
}

// GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetPrimaryMacAddressOk() (*BriefMACAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMacAddress.Get(), o.PrimaryMacAddress.IsSet()
}

// HasPrimaryMacAddress returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasPrimaryMacAddress() bool {
	if o != nil && o.PrimaryMacAddress.IsSet() {
		return true
	}

	return false
}

// SetPrimaryMacAddress gets a reference to the given NullableBriefMACAddressRequest and assigns it to the PrimaryMacAddress field.
func (o *PatchedWritableVMInterfaceRequest) SetPrimaryMacAddress(v BriefMACAddressRequest) {
	o.PrimaryMacAddress.Set(&v)
}

// SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetPrimaryMacAddressNil() {
	o.PrimaryMacAddress.Set(nil)
}

// UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetPrimaryMacAddress() {
	o.PrimaryMacAddress.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableVMInterfaceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetMode() PatchedWritableInterfaceRequestMode {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret PatchedWritableInterfaceRequestMode
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullablePatchedWritableInterfaceRequestMode and assigns it to the Mode field.
func (o *PatchedWritableVMInterfaceRequest) SetMode(v PatchedWritableInterfaceRequestMode) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetMode() {
	o.Mode.Unset()
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetUntaggedVlan() BriefVLANRequest {
	if o == nil || IsNil(o.UntaggedVlan.Get()) {
		var ret BriefVLANRequest
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetUntaggedVlanOk() (*BriefVLANRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableBriefVLANRequest and assigns it to the UntaggedVlan field.
func (o *PatchedWritableVMInterfaceRequest) SetUntaggedVlan(v BriefVLANRequest) {
	o.UntaggedVlan.Set(&v)
}

// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetTaggedVlans() []int32 {
	if o == nil || IsNil(o.TaggedVlans) {
		var ret []int32
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetTaggedVlansOk() ([]int32, bool) {
	if o == nil || IsNil(o.TaggedVlans) {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasTaggedVlans() bool {
	if o != nil && !IsNil(o.TaggedVlans) {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []int32 and assigns it to the TaggedVlans field.
func (o *PatchedWritableVMInterfaceRequest) SetTaggedVlans(v []int32) {
	o.TaggedVlans = v
}

// GetQinqSvlan returns the QinqSvlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetQinqSvlan() BriefVLANRequest {
	if o == nil || IsNil(o.QinqSvlan.Get()) {
		var ret BriefVLANRequest
		return ret
	}
	return *o.QinqSvlan.Get()
}

// GetQinqSvlanOk returns a tuple with the QinqSvlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetQinqSvlanOk() (*BriefVLANRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.QinqSvlan.Get(), o.QinqSvlan.IsSet()
}

// HasQinqSvlan returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasQinqSvlan() bool {
	if o != nil && o.QinqSvlan.IsSet() {
		return true
	}

	return false
}

// SetQinqSvlan gets a reference to the given NullableBriefVLANRequest and assigns it to the QinqSvlan field.
func (o *PatchedWritableVMInterfaceRequest) SetQinqSvlan(v BriefVLANRequest) {
	o.QinqSvlan.Set(&v)
}

// SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetQinqSvlanNil() {
	o.QinqSvlan.Set(nil)
}

// UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetQinqSvlan() {
	o.QinqSvlan.Unset()
}

// GetVlanTranslationPolicy returns the VlanTranslationPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetVlanTranslationPolicy() BriefVLANTranslationPolicyRequest {
	if o == nil || IsNil(o.VlanTranslationPolicy.Get()) {
		var ret BriefVLANTranslationPolicyRequest
		return ret
	}
	return *o.VlanTranslationPolicy.Get()
}

// GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetVlanTranslationPolicyOk() (*BriefVLANTranslationPolicyRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanTranslationPolicy.Get(), o.VlanTranslationPolicy.IsSet()
}

// HasVlanTranslationPolicy returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasVlanTranslationPolicy() bool {
	if o != nil && o.VlanTranslationPolicy.IsSet() {
		return true
	}

	return false
}

// SetVlanTranslationPolicy gets a reference to the given NullableBriefVLANTranslationPolicyRequest and assigns it to the VlanTranslationPolicy field.
func (o *PatchedWritableVMInterfaceRequest) SetVlanTranslationPolicy(v BriefVLANTranslationPolicyRequest) {
	o.VlanTranslationPolicy.Set(&v)
}

// SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetVlanTranslationPolicyNil() {
	o.VlanTranslationPolicy.Set(nil)
}

// UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetVlanTranslationPolicy() {
	o.VlanTranslationPolicy.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVMInterfaceRequest) GetVrf() BriefVRFRequest {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRFRequest
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVMInterfaceRequest) GetVrfOk() (*BriefVRFRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRFRequest and assigns it to the Vrf field.
func (o *PatchedWritableVMInterfaceRequest) SetVrf(v BriefVRFRequest) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *PatchedWritableVMInterfaceRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *PatchedWritableVMInterfaceRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableVMInterfaceRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableVMInterfaceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVMInterfaceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableVMInterfaceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableVMInterfaceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableVMInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VirtualMachine) {
		toSerialize["virtual_machine"] = o.VirtualMachine
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.PrimaryMacAddress.IsSet() {
		toSerialize["primary_mac_address"] = o.PrimaryMacAddress.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if !IsNil(o.TaggedVlans) {
		toSerialize["tagged_vlans"] = o.TaggedVlans
	}
	if o.QinqSvlan.IsSet() {
		toSerialize["qinq_svlan"] = o.QinqSvlan.Get()
	}
	if o.VlanTranslationPolicy.IsSet() {
		toSerialize["vlan_translation_policy"] = o.VlanTranslationPolicy.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
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

func (o *PatchedWritableVMInterfaceRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableVMInterfaceRequest := _PatchedWritableVMInterfaceRequest{}

	err = json.Unmarshal(data, &varPatchedWritableVMInterfaceRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableVMInterfaceRequest(varPatchedWritableVMInterfaceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "primary_mac_address")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "untagged_vlan")
		delete(additionalProperties, "tagged_vlans")
		delete(additionalProperties, "qinq_svlan")
		delete(additionalProperties, "vlan_translation_policy")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableVMInterfaceRequest struct {
	value *PatchedWritableVMInterfaceRequest
	isSet bool
}

func (v NullablePatchedWritableVMInterfaceRequest) Get() *PatchedWritableVMInterfaceRequest {
	return v.value
}

func (v *NullablePatchedWritableVMInterfaceRequest) Set(val *PatchedWritableVMInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVMInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVMInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVMInterfaceRequest(val *PatchedWritableVMInterfaceRequest) *NullablePatchedWritableVMInterfaceRequest {
	return &NullablePatchedWritableVMInterfaceRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVMInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
