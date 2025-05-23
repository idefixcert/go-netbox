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

// IPRangeStatusLabel the model 'IPRangeStatusLabel'
type IPRangeStatusLabel string

// List of IPRange_status_label
const (
	IPRANGESTATUSLABEL_ACTIVE     IPRangeStatusLabel = "Active"
	IPRANGESTATUSLABEL_RESERVED   IPRangeStatusLabel = "Reserved"
	IPRANGESTATUSLABEL_DEPRECATED IPRangeStatusLabel = "Deprecated"
)

// All allowed values of IPRangeStatusLabel enum
var AllowedIPRangeStatusLabelEnumValues = []IPRangeStatusLabel{
	"Active",
	"Reserved",
	"Deprecated",
}

func (v *IPRangeStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPRangeStatusLabel(value)
	for _, existing := range AllowedIPRangeStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPRangeStatusLabel", value)
}

// NewIPRangeStatusLabelFromValue returns a pointer to a valid IPRangeStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPRangeStatusLabelFromValue(v string) (*IPRangeStatusLabel, error) {
	ev := IPRangeStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPRangeStatusLabel: valid values are %v", v, AllowedIPRangeStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPRangeStatusLabel) IsValid() bool {
	for _, existing := range AllowedIPRangeStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPRange_status_label value
func (v IPRangeStatusLabel) Ptr() *IPRangeStatusLabel {
	return &v
}

type NullableIPRangeStatusLabel struct {
	value *IPRangeStatusLabel
	isSet bool
}

func (v NullableIPRangeStatusLabel) Get() *IPRangeStatusLabel {
	return v.value
}

func (v *NullableIPRangeStatusLabel) Set(val *IPRangeStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRangeStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRangeStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRangeStatusLabel(val *IPRangeStatusLabel) *NullableIPRangeStatusLabel {
	return &NullableIPRangeStatusLabel{value: val, isSet: true}
}

func (v NullableIPRangeStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRangeStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
