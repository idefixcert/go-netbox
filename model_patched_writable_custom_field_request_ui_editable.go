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

// PatchedWritableCustomFieldRequestUiEditable Specifies whether the custom field value can be edited in the UI  * `yes` - Yes * `no` - No * `hidden` - Hidden
type PatchedWritableCustomFieldRequestUiEditable string

// List of PatchedWritableCustomFieldRequest_ui_editable
const (
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIEDITABLE_YES    PatchedWritableCustomFieldRequestUiEditable = "yes"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIEDITABLE_NO     PatchedWritableCustomFieldRequestUiEditable = "no"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIEDITABLE_HIDDEN PatchedWritableCustomFieldRequestUiEditable = "hidden"
)

// All allowed values of PatchedWritableCustomFieldRequestUiEditable enum
var AllowedPatchedWritableCustomFieldRequestUiEditableEnumValues = []PatchedWritableCustomFieldRequestUiEditable{
	"yes",
	"no",
	"hidden",
}

func (v *PatchedWritableCustomFieldRequestUiEditable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldRequestUiEditable(value)
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiEditableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldRequestUiEditable", value)
}

// NewPatchedWritableCustomFieldRequestUiEditableFromValue returns a pointer to a valid PatchedWritableCustomFieldRequestUiEditable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldRequestUiEditableFromValue(v string) (*PatchedWritableCustomFieldRequestUiEditable, error) {
	ev := PatchedWritableCustomFieldRequestUiEditable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldRequestUiEditable: valid values are %v", v, AllowedPatchedWritableCustomFieldRequestUiEditableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldRequestUiEditable) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiEditableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldRequest_ui_editable value
func (v PatchedWritableCustomFieldRequestUiEditable) Ptr() *PatchedWritableCustomFieldRequestUiEditable {
	return &v
}

type NullablePatchedWritableCustomFieldRequestUiEditable struct {
	value *PatchedWritableCustomFieldRequestUiEditable
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequestUiEditable) Get() *PatchedWritableCustomFieldRequestUiEditable {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequestUiEditable) Set(val *PatchedWritableCustomFieldRequestUiEditable) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequestUiEditable) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequestUiEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequestUiEditable(val *PatchedWritableCustomFieldRequestUiEditable) *NullablePatchedWritableCustomFieldRequestUiEditable {
	return &NullablePatchedWritableCustomFieldRequestUiEditable{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequestUiEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequestUiEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
