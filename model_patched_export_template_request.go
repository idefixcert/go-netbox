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

// checks if the PatchedExportTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedExportTemplateRequest{}

// PatchedExportTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedExportTemplateRequest struct {
	ObjectTypes []string `json:"object_types,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	// Jinja2 template code. The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode *string `json:"template_code,omitempty"`
	// Defaults to <code>text/plain; charset=utf-8</code>
	MimeType *string `json:"mime_type,omitempty"`
	// Extension to append to the rendered filename
	FileExtension *string `json:"file_extension,omitempty"`
	// Download file as attachment
	AsAttachment         *bool                   `json:"as_attachment,omitempty"`
	DataSource           *BriefDataSourceRequest `json:"data_source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedExportTemplateRequest PatchedExportTemplateRequest

// NewPatchedExportTemplateRequest instantiates a new PatchedExportTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExportTemplateRequest() *PatchedExportTemplateRequest {
	this := PatchedExportTemplateRequest{}
	return &this
}

// NewPatchedExportTemplateRequestWithDefaults instantiates a new PatchedExportTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExportTemplateRequestWithDefaults() *PatchedExportTemplateRequest {
	this := PatchedExportTemplateRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetObjectTypes() []string {
	if o == nil || IsNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasObjectTypes() bool {
	if o != nil && !IsNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedExportTemplateRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedExportTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedExportTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateCode returns the TemplateCode field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetTemplateCode() string {
	if o == nil || IsNil(o.TemplateCode) {
		var ret string
		return ret
	}
	return *o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateCode) {
		return nil, false
	}
	return o.TemplateCode, true
}

// HasTemplateCode returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasTemplateCode() bool {
	if o != nil && !IsNil(o.TemplateCode) {
		return true
	}

	return false
}

// SetTemplateCode gets a reference to the given string and assigns it to the TemplateCode field.
func (o *PatchedExportTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *PatchedExportTemplateRequest) SetMimeType(v string) {
	o.MimeType = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *PatchedExportTemplateRequest) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetAsAttachment returns the AsAttachment field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetAsAttachment() bool {
	if o == nil || IsNil(o.AsAttachment) {
		var ret bool
		return ret
	}
	return *o.AsAttachment
}

// GetAsAttachmentOk returns a tuple with the AsAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetAsAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AsAttachment) {
		return nil, false
	}
	return o.AsAttachment, true
}

// HasAsAttachment returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasAsAttachment() bool {
	if o != nil && !IsNil(o.AsAttachment) {
		return true
	}

	return false
}

// SetAsAttachment gets a reference to the given bool and assigns it to the AsAttachment field.
func (o *PatchedExportTemplateRequest) SetAsAttachment(v bool) {
	o.AsAttachment = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *PatchedExportTemplateRequest) GetDataSource() BriefDataSourceRequest {
	if o == nil || IsNil(o.DataSource) {
		var ret BriefDataSourceRequest
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExportTemplateRequest) GetDataSourceOk() (*BriefDataSourceRequest, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *PatchedExportTemplateRequest) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given BriefDataSourceRequest and assigns it to the DataSource field.
func (o *PatchedExportTemplateRequest) SetDataSource(v BriefDataSourceRequest) {
	o.DataSource = &v
}

func (o PatchedExportTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedExportTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectTypes) {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TemplateCode) {
		toSerialize["template_code"] = o.TemplateCode
	}
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.FileExtension) {
		toSerialize["file_extension"] = o.FileExtension
	}
	if !IsNil(o.AsAttachment) {
		toSerialize["as_attachment"] = o.AsAttachment
	}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedExportTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedExportTemplateRequest := _PatchedExportTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedExportTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedExportTemplateRequest(varPatchedExportTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "file_extension")
		delete(additionalProperties, "as_attachment")
		delete(additionalProperties, "data_source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedExportTemplateRequest struct {
	value *PatchedExportTemplateRequest
	isSet bool
}

func (v NullablePatchedExportTemplateRequest) Get() *PatchedExportTemplateRequest {
	return v.value
}

func (v *NullablePatchedExportTemplateRequest) Set(val *PatchedExportTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExportTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExportTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExportTemplateRequest(val *PatchedExportTemplateRequest) *NullablePatchedExportTemplateRequest {
	return &NullablePatchedExportTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedExportTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExportTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
