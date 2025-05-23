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

// checks if the AggregateFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregateFamily{}

// AggregateFamily struct for AggregateFamily
type AggregateFamily struct {
	Value                *AggregateFamilyValue `json:"value,omitempty"`
	Label                *AggregateFamilyLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregateFamily AggregateFamily

// NewAggregateFamily instantiates a new AggregateFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateFamily() *AggregateFamily {
	this := AggregateFamily{}
	return &this
}

// NewAggregateFamilyWithDefaults instantiates a new AggregateFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateFamilyWithDefaults() *AggregateFamily {
	this := AggregateFamily{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AggregateFamily) GetValue() AggregateFamilyValue {
	if o == nil || IsNil(o.Value) {
		var ret AggregateFamilyValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateFamily) GetValueOk() (*AggregateFamilyValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AggregateFamily) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given AggregateFamilyValue and assigns it to the Value field.
func (o *AggregateFamily) SetValue(v AggregateFamilyValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AggregateFamily) GetLabel() AggregateFamilyLabel {
	if o == nil || IsNil(o.Label) {
		var ret AggregateFamilyLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateFamily) GetLabelOk() (*AggregateFamilyLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AggregateFamily) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given AggregateFamilyLabel and assigns it to the Label field.
func (o *AggregateFamily) SetLabel(v AggregateFamilyLabel) {
	o.Label = &v
}

func (o AggregateFamily) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregateFamily) UnmarshalJSON(data []byte) (err error) {
	varAggregateFamily := _AggregateFamily{}

	err = json.Unmarshal(data, &varAggregateFamily)

	if err != nil {
		return err
	}

	*o = AggregateFamily(varAggregateFamily)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregateFamily struct {
	value *AggregateFamily
	isSet bool
}

func (v NullableAggregateFamily) Get() *AggregateFamily {
	return v.value
}

func (v *NullableAggregateFamily) Set(val *AggregateFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateFamily(val *AggregateFamily) *NullableAggregateFamily {
	return &NullableAggregateFamily{value: val, isSet: true}
}

func (v NullableAggregateFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
