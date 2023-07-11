/*
Aifs api

aifs api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aifsclientgo

import (
	"encoding/json"
)

// checks if the CreateDataViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataViewRequest{}

// CreateDataViewRequest struct for CreateDataViewRequest
type CreateDataViewRequest struct {
	// the name of the data view
	DataViewName *string `json:"dataViewName,omitempty"`
	// the description of the data view
	Description *string `json:"description,omitempty"`
	ViewType *DataViewType `json:"viewType,omitempty"`
	RawDataType *RawDataType `json:"rawDataType,omitempty"`
	ZipFormat *ZipFormat `json:"zipFormat,omitempty"`
	// If it is an annotation type data view, it must have a related raw-data data view
	RelatedDataViewId *string `json:"relatedDataViewId,omitempty"`
	// If it is an annotation type data view, it must have a related annotation template id. If it is a dataset-zip data view, it can have an annotation template id to indicate the annotation template of the annotation data.
	AnnotationTemplateId *string `json:"annotationTemplateId,omitempty"`
	// If it is a dataset-zip type data view, it can have a raw data view id to upload raw data to the data view
	RawDataViewId *string `json:"rawDataViewId,omitempty"`
	// If it is a dataset-zip type data view, it can have a annotation view id to upload annotation data to the data view
	AnnotationViewId *string `json:"annotationViewId,omitempty"`
}

// NewCreateDataViewRequest instantiates a new CreateDataViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataViewRequest() *CreateDataViewRequest {
	this := CreateDataViewRequest{}
	return &this
}

// NewCreateDataViewRequestWithDefaults instantiates a new CreateDataViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataViewRequestWithDefaults() *CreateDataViewRequest {
	this := CreateDataViewRequest{}
	return &this
}

// GetDataViewName returns the DataViewName field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetDataViewName() string {
	if o == nil || IsNil(o.DataViewName) {
		var ret string
		return ret
	}
	return *o.DataViewName
}

// GetDataViewNameOk returns a tuple with the DataViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetDataViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewName) {
		return nil, false
	}
	return o.DataViewName, true
}

// HasDataViewName returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasDataViewName() bool {
	if o != nil && !IsNil(o.DataViewName) {
		return true
	}

	return false
}

// SetDataViewName gets a reference to the given string and assigns it to the DataViewName field.
func (o *CreateDataViewRequest) SetDataViewName(v string) {
	o.DataViewName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDataViewRequest) SetDescription(v string) {
	o.Description = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetViewType() DataViewType {
	if o == nil || IsNil(o.ViewType) {
		var ret DataViewType
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetViewTypeOk() (*DataViewType, bool) {
	if o == nil || IsNil(o.ViewType) {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasViewType() bool {
	if o != nil && !IsNil(o.ViewType) {
		return true
	}

	return false
}

// SetViewType gets a reference to the given DataViewType and assigns it to the ViewType field.
func (o *CreateDataViewRequest) SetViewType(v DataViewType) {
	o.ViewType = &v
}

// GetRawDataType returns the RawDataType field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetRawDataType() RawDataType {
	if o == nil || IsNil(o.RawDataType) {
		var ret RawDataType
		return ret
	}
	return *o.RawDataType
}

// GetRawDataTypeOk returns a tuple with the RawDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetRawDataTypeOk() (*RawDataType, bool) {
	if o == nil || IsNil(o.RawDataType) {
		return nil, false
	}
	return o.RawDataType, true
}

// HasRawDataType returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasRawDataType() bool {
	if o != nil && !IsNil(o.RawDataType) {
		return true
	}

	return false
}

// SetRawDataType gets a reference to the given RawDataType and assigns it to the RawDataType field.
func (o *CreateDataViewRequest) SetRawDataType(v RawDataType) {
	o.RawDataType = &v
}

// GetZipFormat returns the ZipFormat field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetZipFormat() ZipFormat {
	if o == nil || IsNil(o.ZipFormat) {
		var ret ZipFormat
		return ret
	}
	return *o.ZipFormat
}

// GetZipFormatOk returns a tuple with the ZipFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetZipFormatOk() (*ZipFormat, bool) {
	if o == nil || IsNil(o.ZipFormat) {
		return nil, false
	}
	return o.ZipFormat, true
}

// HasZipFormat returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasZipFormat() bool {
	if o != nil && !IsNil(o.ZipFormat) {
		return true
	}

	return false
}

// SetZipFormat gets a reference to the given ZipFormat and assigns it to the ZipFormat field.
func (o *CreateDataViewRequest) SetZipFormat(v ZipFormat) {
	o.ZipFormat = &v
}

// GetRelatedDataViewId returns the RelatedDataViewId field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetRelatedDataViewId() string {
	if o == nil || IsNil(o.RelatedDataViewId) {
		var ret string
		return ret
	}
	return *o.RelatedDataViewId
}

// GetRelatedDataViewIdOk returns a tuple with the RelatedDataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetRelatedDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedDataViewId) {
		return nil, false
	}
	return o.RelatedDataViewId, true
}

// HasRelatedDataViewId returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasRelatedDataViewId() bool {
	if o != nil && !IsNil(o.RelatedDataViewId) {
		return true
	}

	return false
}

// SetRelatedDataViewId gets a reference to the given string and assigns it to the RelatedDataViewId field.
func (o *CreateDataViewRequest) SetRelatedDataViewId(v string) {
	o.RelatedDataViewId = &v
}

// GetAnnotationTemplateId returns the AnnotationTemplateId field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetAnnotationTemplateId() string {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		var ret string
		return ret
	}
	return *o.AnnotationTemplateId
}

// GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetAnnotationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		return nil, false
	}
	return o.AnnotationTemplateId, true
}

// HasAnnotationTemplateId returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasAnnotationTemplateId() bool {
	if o != nil && !IsNil(o.AnnotationTemplateId) {
		return true
	}

	return false
}

// SetAnnotationTemplateId gets a reference to the given string and assigns it to the AnnotationTemplateId field.
func (o *CreateDataViewRequest) SetAnnotationTemplateId(v string) {
	o.AnnotationTemplateId = &v
}

// GetRawDataViewId returns the RawDataViewId field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetRawDataViewId() string {
	if o == nil || IsNil(o.RawDataViewId) {
		var ret string
		return ret
	}
	return *o.RawDataViewId
}

// GetRawDataViewIdOk returns a tuple with the RawDataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetRawDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.RawDataViewId) {
		return nil, false
	}
	return o.RawDataViewId, true
}

// HasRawDataViewId returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasRawDataViewId() bool {
	if o != nil && !IsNil(o.RawDataViewId) {
		return true
	}

	return false
}

// SetRawDataViewId gets a reference to the given string and assigns it to the RawDataViewId field.
func (o *CreateDataViewRequest) SetRawDataViewId(v string) {
	o.RawDataViewId = &v
}

// GetAnnotationViewId returns the AnnotationViewId field value if set, zero value otherwise.
func (o *CreateDataViewRequest) GetAnnotationViewId() string {
	if o == nil || IsNil(o.AnnotationViewId) {
		var ret string
		return ret
	}
	return *o.AnnotationViewId
}

// GetAnnotationViewIdOk returns a tuple with the AnnotationViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewRequest) GetAnnotationViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationViewId) {
		return nil, false
	}
	return o.AnnotationViewId, true
}

// HasAnnotationViewId returns a boolean if a field has been set.
func (o *CreateDataViewRequest) HasAnnotationViewId() bool {
	if o != nil && !IsNil(o.AnnotationViewId) {
		return true
	}

	return false
}

// SetAnnotationViewId gets a reference to the given string and assigns it to the AnnotationViewId field.
func (o *CreateDataViewRequest) SetAnnotationViewId(v string) {
	o.AnnotationViewId = &v
}

func (o CreateDataViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataViewName) {
		toSerialize["dataViewName"] = o.DataViewName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ViewType) {
		toSerialize["viewType"] = o.ViewType
	}
	if !IsNil(o.RawDataType) {
		toSerialize["rawDataType"] = o.RawDataType
	}
	if !IsNil(o.ZipFormat) {
		toSerialize["zipFormat"] = o.ZipFormat
	}
	if !IsNil(o.RelatedDataViewId) {
		toSerialize["relatedDataViewId"] = o.RelatedDataViewId
	}
	if !IsNil(o.AnnotationTemplateId) {
		toSerialize["annotationTemplateId"] = o.AnnotationTemplateId
	}
	if !IsNil(o.RawDataViewId) {
		toSerialize["rawDataViewId"] = o.RawDataViewId
	}
	if !IsNil(o.AnnotationViewId) {
		toSerialize["annotationViewId"] = o.AnnotationViewId
	}
	return toSerialize, nil
}

type NullableCreateDataViewRequest struct {
	value *CreateDataViewRequest
	isSet bool
}

func (v NullableCreateDataViewRequest) Get() *CreateDataViewRequest {
	return v.value
}

func (v *NullableCreateDataViewRequest) Set(val *CreateDataViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataViewRequest(val *CreateDataViewRequest) *NullableCreateDataViewRequest {
	return &NullableCreateDataViewRequest{value: val, isSet: true}
}

func (v NullableCreateDataViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


