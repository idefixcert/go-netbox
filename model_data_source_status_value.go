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

// DataSourceStatusValue * `new` - New * `queued` - Queued * `syncing` - Syncing * `completed` - Completed * `failed` - Failed
type DataSourceStatusValue string

// List of DataSource_status_value
const (
	DATASOURCESTATUSVALUE_NEW       DataSourceStatusValue = "new"
	DATASOURCESTATUSVALUE_QUEUED    DataSourceStatusValue = "queued"
	DATASOURCESTATUSVALUE_SYNCING   DataSourceStatusValue = "syncing"
	DATASOURCESTATUSVALUE_COMPLETED DataSourceStatusValue = "completed"
	DATASOURCESTATUSVALUE_FAILED    DataSourceStatusValue = "failed"
)

// All allowed values of DataSourceStatusValue enum
var AllowedDataSourceStatusValueEnumValues = []DataSourceStatusValue{
	"new",
	"queued",
	"syncing",
	"completed",
	"failed",
}

func (v *DataSourceStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceStatusValue(value)
	for _, existing := range AllowedDataSourceStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceStatusValue", value)
}

// NewDataSourceStatusValueFromValue returns a pointer to a valid DataSourceStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceStatusValueFromValue(v string) (*DataSourceStatusValue, error) {
	ev := DataSourceStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceStatusValue: valid values are %v", v, AllowedDataSourceStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceStatusValue) IsValid() bool {
	for _, existing := range AllowedDataSourceStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSource_status_value value
func (v DataSourceStatusValue) Ptr() *DataSourceStatusValue {
	return &v
}

type NullableDataSourceStatusValue struct {
	value *DataSourceStatusValue
	isSet bool
}

func (v NullableDataSourceStatusValue) Get() *DataSourceStatusValue {
	return v.value
}

func (v *NullableDataSourceStatusValue) Set(val *DataSourceStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceStatusValue(val *DataSourceStatusValue) *NullableDataSourceStatusValue {
	return &NullableDataSourceStatusValue{value: val, isSet: true}
}

func (v NullableDataSourceStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
