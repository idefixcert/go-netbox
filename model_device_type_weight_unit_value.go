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

// DeviceTypeWeightUnitValue * `kg` - Kilograms * `g` - Grams * `lb` - Pounds * `oz` - Ounces
type DeviceTypeWeightUnitValue string

// List of DeviceType_weight_unit_value
const (
	DEVICETYPEWEIGHTUNITVALUE_KG    DeviceTypeWeightUnitValue = "kg"
	DEVICETYPEWEIGHTUNITVALUE_G     DeviceTypeWeightUnitValue = "g"
	DEVICETYPEWEIGHTUNITVALUE_LB    DeviceTypeWeightUnitValue = "lb"
	DEVICETYPEWEIGHTUNITVALUE_OZ    DeviceTypeWeightUnitValue = "oz"
	DEVICETYPEWEIGHTUNITVALUE_EMPTY DeviceTypeWeightUnitValue = ""
)

// All allowed values of DeviceTypeWeightUnitValue enum
var AllowedDeviceTypeWeightUnitValueEnumValues = []DeviceTypeWeightUnitValue{
	"kg",
	"g",
	"lb",
	"oz",
	"",
}

func (v *DeviceTypeWeightUnitValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeWeightUnitValue(value)
	for _, existing := range AllowedDeviceTypeWeightUnitValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeWeightUnitValue", value)
}

// NewDeviceTypeWeightUnitValueFromValue returns a pointer to a valid DeviceTypeWeightUnitValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeWeightUnitValueFromValue(v string) (*DeviceTypeWeightUnitValue, error) {
	ev := DeviceTypeWeightUnitValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeWeightUnitValue: valid values are %v", v, AllowedDeviceTypeWeightUnitValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeWeightUnitValue) IsValid() bool {
	for _, existing := range AllowedDeviceTypeWeightUnitValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceType_weight_unit_value value
func (v DeviceTypeWeightUnitValue) Ptr() *DeviceTypeWeightUnitValue {
	return &v
}

type NullableDeviceTypeWeightUnitValue struct {
	value *DeviceTypeWeightUnitValue
	isSet bool
}

func (v NullableDeviceTypeWeightUnitValue) Get() *DeviceTypeWeightUnitValue {
	return v.value
}

func (v *NullableDeviceTypeWeightUnitValue) Set(val *DeviceTypeWeightUnitValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeWeightUnitValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeWeightUnitValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeWeightUnitValue(val *DeviceTypeWeightUnitValue) *NullableDeviceTypeWeightUnitValue {
	return &NullableDeviceTypeWeightUnitValue{value: val, isSet: true}
}

func (v NullableDeviceTypeWeightUnitValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeWeightUnitValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
