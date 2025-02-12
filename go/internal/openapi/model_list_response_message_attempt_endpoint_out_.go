/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListResponseMessageAttemptEndpointOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseMessageAttemptEndpointOut{}

// ListResponseMessageAttemptEndpointOut struct for ListResponseMessageAttemptEndpointOut
type ListResponseMessageAttemptEndpointOut struct {
	Data []MessageAttemptEndpointOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseMessageAttemptEndpointOut ListResponseMessageAttemptEndpointOut

// NewListResponseMessageAttemptEndpointOut instantiates a new ListResponseMessageAttemptEndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseMessageAttemptEndpointOut(data []MessageAttemptEndpointOut, done bool, iterator NullableString) *ListResponseMessageAttemptEndpointOut {
	this := ListResponseMessageAttemptEndpointOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewListResponseMessageAttemptEndpointOutWithDefaults instantiates a new ListResponseMessageAttemptEndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseMessageAttemptEndpointOutWithDefaults() *ListResponseMessageAttemptEndpointOut {
	this := ListResponseMessageAttemptEndpointOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseMessageAttemptEndpointOut) GetData() []MessageAttemptEndpointOut {
	if o == nil {
		var ret []MessageAttemptEndpointOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseMessageAttemptEndpointOut) GetDataOk() ([]MessageAttemptEndpointOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseMessageAttemptEndpointOut) SetData(v []MessageAttemptEndpointOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseMessageAttemptEndpointOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseMessageAttemptEndpointOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseMessageAttemptEndpointOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListResponseMessageAttemptEndpointOut) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseMessageAttemptEndpointOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// SetIterator sets field value
func (o *ListResponseMessageAttemptEndpointOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseMessageAttemptEndpointOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseMessageAttemptEndpointOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseMessageAttemptEndpointOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseMessageAttemptEndpointOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseMessageAttemptEndpointOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseMessageAttemptEndpointOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseMessageAttemptEndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseMessageAttemptEndpointOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator.Get()
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseMessageAttemptEndpointOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
		"iterator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListResponseMessageAttemptEndpointOut := _ListResponseMessageAttemptEndpointOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseMessageAttemptEndpointOut)

	if err != nil {
		return err
	}

	*o = ListResponseMessageAttemptEndpointOut(varListResponseMessageAttemptEndpointOut)

	return err
}

type NullableListResponseMessageAttemptEndpointOut struct {
	value *ListResponseMessageAttemptEndpointOut
	isSet bool
}

func (v NullableListResponseMessageAttemptEndpointOut) Get() *ListResponseMessageAttemptEndpointOut {
	return v.value
}

func (v *NullableListResponseMessageAttemptEndpointOut) Set(val *ListResponseMessageAttemptEndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseMessageAttemptEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseMessageAttemptEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseMessageAttemptEndpointOut(val *ListResponseMessageAttemptEndpointOut) *NullableListResponseMessageAttemptEndpointOut {
	return &NullableListResponseMessageAttemptEndpointOut{value: val, isSet: true}
}

func (v NullableListResponseMessageAttemptEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseMessageAttemptEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


