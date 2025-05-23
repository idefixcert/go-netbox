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

// EventRuleActionTypeValue * `webhook` - Webhook * `script` - Script * `notification` - Notification
type EventRuleActionTypeValue string

// List of EventRule_action_type_value
const (
	EVENTRULEACTIONTYPEVALUE_WEBHOOK      EventRuleActionTypeValue = "webhook"
	EVENTRULEACTIONTYPEVALUE_SCRIPT       EventRuleActionTypeValue = "script"
	EVENTRULEACTIONTYPEVALUE_NOTIFICATION EventRuleActionTypeValue = "notification"
)

// All allowed values of EventRuleActionTypeValue enum
var AllowedEventRuleActionTypeValueEnumValues = []EventRuleActionTypeValue{
	"webhook",
	"script",
	"notification",
}

func (v *EventRuleActionTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventRuleActionTypeValue(value)
	for _, existing := range AllowedEventRuleActionTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventRuleActionTypeValue", value)
}

// NewEventRuleActionTypeValueFromValue returns a pointer to a valid EventRuleActionTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventRuleActionTypeValueFromValue(v string) (*EventRuleActionTypeValue, error) {
	ev := EventRuleActionTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventRuleActionTypeValue: valid values are %v", v, AllowedEventRuleActionTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventRuleActionTypeValue) IsValid() bool {
	for _, existing := range AllowedEventRuleActionTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventRule_action_type_value value
func (v EventRuleActionTypeValue) Ptr() *EventRuleActionTypeValue {
	return &v
}

type NullableEventRuleActionTypeValue struct {
	value *EventRuleActionTypeValue
	isSet bool
}

func (v NullableEventRuleActionTypeValue) Get() *EventRuleActionTypeValue {
	return v.value
}

func (v *NullableEventRuleActionTypeValue) Set(val *EventRuleActionTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRuleActionTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRuleActionTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRuleActionTypeValue(val *EventRuleActionTypeValue) *NullableEventRuleActionTypeValue {
	return &NullableEventRuleActionTypeValue{value: val, isSet: true}
}

func (v NullableEventRuleActionTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRuleActionTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
