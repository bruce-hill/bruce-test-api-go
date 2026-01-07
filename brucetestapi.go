// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/bruce-hill/bruce-test-api-go/internal/apiform"
	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	shimjson "github.com/bruce-hill/bruce-test-api-go/internal/encoding/json"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
)

func SomethingParamOfSomethingObject(name string) SomethingUnionParam {
	var variant SomethingObjectParam
	variant.Name = name
	return SomethingUnionParam{OfSomethingObject: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SomethingUnionParam struct {
	OfFloat           param.Opt[float64]    `json:",omitzero,inline"`
	OfSomethingObject *SomethingObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u SomethingUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfSomethingObject)
}
func (u *SomethingUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SomethingUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfSomethingObject) {
		return u.OfSomethingObject
	}
	return nil
}

// The property Name is required.
type SomethingObjectParam struct {
	Name  string           `json:"name,required"`
	Count param.Opt[int64] `json:"count,omitzero"`
	paramObj
}

func (r SomethingObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow SomethingObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SomethingObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormTestResponse = any

type JsonTestResponse = any

type FormTestParams struct {
	Version     int64                        `path:"version,required" json:"-"`
	Date        time.Time                    `query:"date,required" format:"date" json:"-"`
	Datetime    time.Time                    `query:"datetime,required" format:"date-time" json:"-"`
	Time        string                       `query:"time,required" format:"time" json:"-"`
	Limit       param.Opt[int64]             `query:"limit,omitzero" json:"-"`
	Blorp       param.Opt[string]            `json:"blorp,omitzero"`
	XTraceID    param.Opt[string]            `header:"X-Trace-ID,omitzero" json:"-"`
	PlsNull     any                          `json:"pls_null"`
	Filter      FormTestParamsFilter         `query:"filter,omitzero" json:"-"`
	IDOrIndex   FormTestParamsIDOrIndexUnion `query:"idOrIndex,omitzero" json:"-"`
	Tags        []string                     `query:"tags,omitzero" json:"-"`
	Pets        []FormTestParamsPet          `json:"pets,omitzero"`
	Preferences FormTestParamsPreferences    `json:"preferences,omitzero"`
	Something   SomethingUnionParam          `json:"something,omitzero"`
	XFlags      []string                     `header:"X-Flags,omitzero" json:"-"`
	paramObj
}

func (r FormTestParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [FormTestParams]'s query parameters as `url.Values`.
func (r FormTestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FormTestParamsFilter struct {
	Status param.Opt[string]        `query:"status,omitzero" json:"-"`
	Meta   FormTestParamsFilterMeta `query:"meta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormTestParamsFilter]'s query parameters as `url.Values`.
func (r FormTestParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FormTestParamsFilterMeta struct {
	Level param.Opt[int64] `query:"level,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormTestParamsFilterMeta]'s query parameters as
// `url.Values`.
func (r FormTestParamsFilterMeta) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FormTestParamsIDOrIndexUnion struct {
	OfInt    param.Opt[int64]  `query:",omitzero,inline"`
	OfString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

func (u *FormTestParamsIDOrIndexUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property Name is required.
type FormTestParamsPet struct {
	Name string           `json:"name,required"`
	Age  param.Opt[int64] `json:"age,omitzero"`
	paramObj
}

func (r FormTestParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormTestParamsPreferences struct {
	Alerts param.Opt[bool]   `json:"alerts,omitzero"`
	Theme  param.Opt[string] `json:"theme,omitzero"`
	paramObj
}

func (r FormTestParamsPreferences) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JsonTestParams struct {
	Version     int64                        `path:"version,required" json:"-"`
	Date        time.Time                    `query:"date,required" format:"date" json:"-"`
	Datetime    time.Time                    `query:"datetime,required" format:"date-time" json:"-"`
	Time        string                       `query:"time,required" format:"time" json:"-"`
	Limit       param.Opt[int64]             `query:"limit,omitzero" json:"-"`
	Blorp       param.Opt[string]            `json:"blorp,omitzero"`
	XTraceID    param.Opt[string]            `header:"X-Trace-ID,omitzero" json:"-"`
	PlsNull     any                          `json:"pls_null"`
	Filter      JsonTestParamsFilter         `query:"filter,omitzero" json:"-"`
	IDOrIndex   JsonTestParamsIDOrIndexUnion `query:"idOrIndex,omitzero" json:"-"`
	Tags        []string                     `query:"tags,omitzero" json:"-"`
	Pets        []JsonTestParamsPet          `json:"pets,omitzero"`
	Preferences JsonTestParamsPreferences    `json:"preferences,omitzero"`
	Something   SomethingUnionParam          `json:"something,omitzero"`
	XFlags      []string                     `header:"X-Flags,omitzero" json:"-"`
	paramObj
}

func (r JsonTestParams) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [JsonTestParams]'s query parameters as `url.Values`.
func (r JsonTestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type JsonTestParamsFilter struct {
	Status param.Opt[string]        `query:"status,omitzero" json:"-"`
	Meta   JsonTestParamsFilterMeta `query:"meta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JsonTestParamsFilter]'s query parameters as `url.Values`.
func (r JsonTestParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type JsonTestParamsFilterMeta struct {
	Level param.Opt[int64] `query:"level,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JsonTestParamsFilterMeta]'s query parameters as
// `url.Values`.
func (r JsonTestParamsFilterMeta) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JsonTestParamsIDOrIndexUnion struct {
	OfInt    param.Opt[int64]  `query:",omitzero,inline"`
	OfString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

func (u *JsonTestParamsIDOrIndexUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property Name is required.
type JsonTestParamsPet struct {
	Name string           `json:"name,required"`
	Age  param.Opt[int64] `json:"age,omitzero"`
	paramObj
}

func (r JsonTestParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JsonTestParamsPreferences struct {
	Alerts param.Opt[bool]   `json:"alerts,omitzero"`
	Theme  param.Opt[string] `json:"theme,omitzero"`
	paramObj
}

func (r JsonTestParamsPreferences) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParamsPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParamsPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateCountParams struct {
	Body int64
	paramObj
}

func (r UpdateCountParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *UpdateCountParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
