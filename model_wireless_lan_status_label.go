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

// WirelessLANStatusLabel the model 'WirelessLANStatusLabel'
type WirelessLANStatusLabel string

// List of WirelessLAN_status_label
const (
	WIRELESSLANSTATUSLABEL_ACTIVE     WirelessLANStatusLabel = "Active"
	WIRELESSLANSTATUSLABEL_RESERVED   WirelessLANStatusLabel = "Reserved"
	WIRELESSLANSTATUSLABEL_DISABLED   WirelessLANStatusLabel = "Disabled"
	WIRELESSLANSTATUSLABEL_DEPRECATED WirelessLANStatusLabel = "Deprecated"
)

// All allowed values of WirelessLANStatusLabel enum
var AllowedWirelessLANStatusLabelEnumValues = []WirelessLANStatusLabel{
	"Active",
	"Reserved",
	"Disabled",
	"Deprecated",
}

func (v *WirelessLANStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLANStatusLabel(value)
	for _, existing := range AllowedWirelessLANStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLANStatusLabel", value)
}

// NewWirelessLANStatusLabelFromValue returns a pointer to a valid WirelessLANStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLANStatusLabelFromValue(v string) (*WirelessLANStatusLabel, error) {
	ev := WirelessLANStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLANStatusLabel: valid values are %v", v, AllowedWirelessLANStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLANStatusLabel) IsValid() bool {
	for _, existing := range AllowedWirelessLANStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLAN_status_label value
func (v WirelessLANStatusLabel) Ptr() *WirelessLANStatusLabel {
	return &v
}

type NullableWirelessLANStatusLabel struct {
	value *WirelessLANStatusLabel
	isSet bool
}

func (v NullableWirelessLANStatusLabel) Get() *WirelessLANStatusLabel {
	return v.value
}

func (v *NullableWirelessLANStatusLabel) Set(val *WirelessLANStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANStatusLabel(val *WirelessLANStatusLabel) *NullableWirelessLANStatusLabel {
	return &NullableWirelessLANStatusLabel{value: val, isSet: true}
}

func (v NullableWirelessLANStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
