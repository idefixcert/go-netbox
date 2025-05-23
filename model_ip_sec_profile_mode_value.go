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

// IPSecProfileModeValue * `esp` - ESP * `ah` - AH
type IPSecProfileModeValue string

// List of IPSecProfile_mode_value
const (
	IPSECPROFILEMODEVALUE_ESP IPSecProfileModeValue = "esp"
	IPSECPROFILEMODEVALUE_AH  IPSecProfileModeValue = "ah"
)

// All allowed values of IPSecProfileModeValue enum
var AllowedIPSecProfileModeValueEnumValues = []IPSecProfileModeValue{
	"esp",
	"ah",
}

func (v *IPSecProfileModeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPSecProfileModeValue(value)
	for _, existing := range AllowedIPSecProfileModeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPSecProfileModeValue", value)
}

// NewIPSecProfileModeValueFromValue returns a pointer to a valid IPSecProfileModeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPSecProfileModeValueFromValue(v string) (*IPSecProfileModeValue, error) {
	ev := IPSecProfileModeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPSecProfileModeValue: valid values are %v", v, AllowedIPSecProfileModeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPSecProfileModeValue) IsValid() bool {
	for _, existing := range AllowedIPSecProfileModeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPSecProfile_mode_value value
func (v IPSecProfileModeValue) Ptr() *IPSecProfileModeValue {
	return &v
}

type NullableIPSecProfileModeValue struct {
	value *IPSecProfileModeValue
	isSet bool
}

func (v NullableIPSecProfileModeValue) Get() *IPSecProfileModeValue {
	return v.value
}

func (v *NullableIPSecProfileModeValue) Set(val *IPSecProfileModeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfileModeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfileModeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfileModeValue(val *IPSecProfileModeValue) *NullableIPSecProfileModeValue {
	return &NullableIPSecProfileModeValue{value: val, isSet: true}
}

func (v NullableIPSecProfileModeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfileModeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
