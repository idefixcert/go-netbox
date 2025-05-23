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

// IPSecProfileModeLabel the model 'IPSecProfileModeLabel'
type IPSecProfileModeLabel string

// List of IPSecProfile_mode_label
const (
	IPSECPROFILEMODELABEL_ESP IPSecProfileModeLabel = "ESP"
	IPSECPROFILEMODELABEL_AH  IPSecProfileModeLabel = "AH"
)

// All allowed values of IPSecProfileModeLabel enum
var AllowedIPSecProfileModeLabelEnumValues = []IPSecProfileModeLabel{
	"ESP",
	"AH",
}

func (v *IPSecProfileModeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPSecProfileModeLabel(value)
	for _, existing := range AllowedIPSecProfileModeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPSecProfileModeLabel", value)
}

// NewIPSecProfileModeLabelFromValue returns a pointer to a valid IPSecProfileModeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPSecProfileModeLabelFromValue(v string) (*IPSecProfileModeLabel, error) {
	ev := IPSecProfileModeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPSecProfileModeLabel: valid values are %v", v, AllowedIPSecProfileModeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPSecProfileModeLabel) IsValid() bool {
	for _, existing := range AllowedIPSecProfileModeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPSecProfile_mode_label value
func (v IPSecProfileModeLabel) Ptr() *IPSecProfileModeLabel {
	return &v
}

type NullableIPSecProfileModeLabel struct {
	value *IPSecProfileModeLabel
	isSet bool
}

func (v NullableIPSecProfileModeLabel) Get() *IPSecProfileModeLabel {
	return v.value
}

func (v *NullableIPSecProfileModeLabel) Set(val *IPSecProfileModeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfileModeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfileModeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfileModeLabel(val *IPSecProfileModeLabel) *NullableIPSecProfileModeLabel {
	return &NullableIPSecProfileModeLabel{value: val, isSet: true}
}

func (v NullableIPSecProfileModeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfileModeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
