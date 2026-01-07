// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	shimjson "github.com/bruce-hill/bruce-test-api-go/internal/encoding/json"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
)

type JsonTestResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	UserID    string    `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JsonTestResponse) RawJSON() string { return r.JSON.raw }
func (r *JsonTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JsonTestParams struct {
	Version        int64                              `path:"version,required" json:"-"`
	Date           time.Time                          `query:"date,required" format:"date" json:"-"`
	Datetime       time.Time                          `query:"datetime,required" format:"date-time" json:"-"`
	Time           string                             `query:"time,required" format:"time" json:"-"`
	Limit          param.Opt[int64]                   `query:"limit,omitzero" json:"-"`
	Blorp          param.Opt[string]                  `json:"blorp,omitzero"`
	XTraceID       param.Opt[string]                  `header:"X-Trace-ID,omitzero" json:"-"`
	PlsNull        any                                `json:"pls_null"`
	Filter         JsonTestParamsFilter               `query:"filter,omitzero" json:"-"`
	IDOrIndex      JsonTestParamsIDOrIndexUnion       `query:"idOrIndex,omitzero" json:"-"`
	Tags           []string                           `query:"tags,omitzero" json:"-"`
	ManySomethings []JsonTestParamsManySomethingUnion `json:"many_somethings,omitzero"`
	Pets           []JsonTestParamsPet                `json:"pets,omitzero"`
	Preferences    JsonTestParamsPreferences          `json:"preferences,omitzero"`
	Something      JsonTestParamsSomethingUnion       `json:"something,omitzero"`
	XFlags         []string                           `header:"X-Flags,omitzero" json:"-"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JsonTestParamsManySomethingUnion struct {
	OfFloat  param.Opt[float64]                 `json:",omitzero,inline"`
	OfThingy *JsonTestParamsManySomethingThingy `json:",omitzero,inline"`
	paramUnion
}

func (u JsonTestParamsManySomethingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfThingy)
}
func (u *JsonTestParamsManySomethingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JsonTestParamsManySomethingUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfThingy) {
		return u.OfThingy
	}
	return nil
}

// The property Name is required.
type JsonTestParamsManySomethingThingy struct {
	Name  string           `json:"name,required"`
	Count param.Opt[int64] `json:"count,omitzero"`
	paramObj
}

func (r JsonTestParamsManySomethingThingy) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParamsManySomethingThingy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParamsManySomethingThingy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JsonTestParamsSomethingUnion struct {
	OfFloat  param.Opt[float64]             `json:",omitzero,inline"`
	OfThingy *JsonTestParamsSomethingThingy `json:",omitzero,inline"`
	paramUnion
}

func (u JsonTestParamsSomethingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfThingy)
}
func (u *JsonTestParamsSomethingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JsonTestParamsSomethingUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfThingy) {
		return u.OfThingy
	}
	return nil
}

// The property Name is required.
type JsonTestParamsSomethingThingy struct {
	Name  string           `json:"name,required"`
	Count param.Opt[int64] `json:"count,omitzero"`
	paramObj
}

func (r JsonTestParamsSomethingThingy) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParamsSomethingThingy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParamsSomethingThingy) UnmarshalJSON(data []byte) error {
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
