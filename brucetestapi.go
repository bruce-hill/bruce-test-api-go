// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/bruce-hill/bruce-test-api-go/internal/apiform"
	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	shimjson "github.com/bruce-hill/bruce-test-api-go/internal/encoding/json"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
)

// Response confirming the user update
type FormTestResponse struct {
	// Status of the update operation
	Status string `json:"status"`
	// Timestamp when the update occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// The updated user's ID
	UserID string `json:"userId"`
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
func (r FormTestResponse) RawJSON() string { return r.JSON.raw }
func (r *FormTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response confirming the user update
type JsonTestResponse struct {
	// Status of the update operation
	Status string `json:"status"`
	// Timestamp when the update occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// The updated user's ID
	UserID string `json:"userId"`
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

type UpdateCountResponse struct {
	Count   int64 `json:"count"`
	Success bool  `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateCountResponse) RawJSON() string { return r.JSON.raw }
func (r *UpdateCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadTestResponse struct {
	// Human-readable status message.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadTestResponse) RawJSON() string { return r.JSON.raw }
func (r *UploadTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormTestParams struct {
	// The API version to use
	Version int64 `path:"version,required" json:"-"`
	// Date filter in ISO 8601 format (YYYY-MM-DD)
	Date time.Time `query:"date,required" format:"date" json:"-"`
	// Full datetime filter in ISO 8601 format
	Datetime time.Time `query:"datetime,required" format:"date-time" json:"-"`
	// Time filter in ISO 8601 format (HH:MM:SS)
	Time string `query:"time,required" format:"time" json:"-"`
	// Required field for demonstration purposes
	Blorp string `json:"blorp,required"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Trace ID string for distributed tracing
	XTraceID param.Opt[string] `header:"X-Trace-ID,omitzero" json:"-"`
	// A null value field for testing null handling
	PlsNull any `json:"pls_null"`
	// Complex filter object for advanced querying
	Filter FormTestParamsFilter `query:"filter,omitzero" json:"-"`
	// Flexible identifier that can be either a numeric index or string ID
	IDOrIndex FormTestParamsIDOrIndexUnion `query:"idOrIndex,omitzero" json:"-"`
	// Filter results by one or more tags
	Tags []string `query:"tags,omitzero" json:"-"`
	// Array of Something items
	ManySomethings []FormTestParamsManySomethingUnion `json:"many_somethings,omitzero"`
	// List of user's pets
	Pets []FormTestParamsPet `json:"pets,omitzero"`
	// User preference settings
	Preferences FormTestParamsPreferences `json:"preferences,omitzero"`
	// A flexible type that can be either a number or an object with name and optional
	// count properties
	Something FormTestParamsSomethingUnion `json:"something,omitzero"`
	// Array of feature flag names
	XFlags []string `header:"X-Flags,omitzero" json:"-"`
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

// Complex filter object for advanced querying
type FormTestParamsFilter struct {
	// Filter by status value
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Metadata filter options
	Meta FormTestParamsFilterMeta `query:"meta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormTestParamsFilter]'s query parameters as `url.Values`.
func (r FormTestParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Metadata filter options
type FormTestParamsFilterMeta struct {
	// Filter by level value
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FormTestParamsManySomethingUnion struct {
	OfFloat  param.Opt[float64]                 `json:",omitzero,inline"`
	OfThingy *FormTestParamsManySomethingThingy `json:",omitzero,inline"`
	paramUnion
}

func (u FormTestParamsManySomethingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfThingy)
}
func (u *FormTestParamsManySomethingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FormTestParamsManySomethingUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfThingy) {
		return u.OfThingy
	}
	return nil
}

// An object with a required name and optional count
//
// The property Name is required.
type FormTestParamsManySomethingThingy struct {
	// Name identifier
	Name string `json:"name,required"`
	// Optional count value
	Count param.Opt[int64] `json:"count,omitzero"`
	paramObj
}

func (r FormTestParamsManySomethingThingy) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsManySomethingThingy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsManySomethingThingy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pet information
//
// The property Name is required.
type FormTestParamsPet struct {
	// Name of the pet
	Name string `json:"name,required"`
	// Age of the pet in years
	Age param.Opt[int64] `json:"age,omitzero"`
	paramObj
}

func (r FormTestParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User preference settings
type FormTestParamsPreferences struct {
	// Whether to enable alert notifications
	Alerts param.Opt[bool] `json:"alerts,omitzero"`
	// UI theme preference (e.g., dark, light)
	Theme param.Opt[string] `json:"theme,omitzero"`
	paramObj
}

func (r FormTestParamsPreferences) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FormTestParamsSomethingUnion struct {
	OfFloat  param.Opt[float64]             `json:",omitzero,inline"`
	OfThingy *FormTestParamsSomethingThingy `json:",omitzero,inline"`
	paramUnion
}

func (u FormTestParamsSomethingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfThingy)
}
func (u *FormTestParamsSomethingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FormTestParamsSomethingUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfThingy) {
		return u.OfThingy
	}
	return nil
}

// An object with a required name and optional count
//
// The property Name is required.
type FormTestParamsSomethingThingy struct {
	// Name identifier
	Name string `json:"name,required"`
	// Optional count value
	Count param.Opt[int64] `json:"count,omitzero"`
	paramObj
}

func (r FormTestParamsSomethingThingy) MarshalJSON() (data []byte, err error) {
	type shadow FormTestParamsSomethingThingy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormTestParamsSomethingThingy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JsonTestParams struct {
	// The API version to use
	Version int64 `path:"version,required" json:"-"`
	// Date filter in ISO 8601 format (YYYY-MM-DD)
	Date time.Time `query:"date,required" format:"date" json:"-"`
	// Full datetime filter in ISO 8601 format
	Datetime time.Time `query:"datetime,required" format:"date-time" json:"-"`
	// Time filter in ISO 8601 format (HH:MM:SS)
	Time string `query:"time,required" format:"time" json:"-"`
	// Required field for demonstration purposes
	Blorp string `json:"blorp,required"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Trace ID string for distributed tracing
	XTraceID param.Opt[string] `header:"X-Trace-ID,omitzero" json:"-"`
	// A null value field for testing null handling
	PlsNull any `json:"pls_null"`
	// Complex filter object for advanced querying
	Filter JsonTestParamsFilter `query:"filter,omitzero" json:"-"`
	// Flexible identifier that can be either a numeric index or string ID
	IDOrIndex JsonTestParamsIDOrIndexUnion `query:"idOrIndex,omitzero" json:"-"`
	// Filter results by one or more tags
	Tags []string `query:"tags,omitzero" json:"-"`
	// Array of Something items
	ManySomethings []JsonTestParamsManySomethingUnion `json:"many_somethings,omitzero"`
	// List of user's pets
	Pets []JsonTestParamsPet `json:"pets,omitzero"`
	// User preference settings
	Preferences JsonTestParamsPreferences `json:"preferences,omitzero"`
	// A flexible type that can be either a number or an object with name and optional
	// count properties
	Something JsonTestParamsSomethingUnion `json:"something,omitzero"`
	// Array of feature flag names
	XFlags []string `header:"X-Flags,omitzero" json:"-"`
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

// Complex filter object for advanced querying
type JsonTestParamsFilter struct {
	// Filter by status value
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Metadata filter options
	Meta JsonTestParamsFilterMeta `query:"meta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JsonTestParamsFilter]'s query parameters as `url.Values`.
func (r JsonTestParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Metadata filter options
type JsonTestParamsFilterMeta struct {
	// Filter by level value
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

// An object with a required name and optional count
//
// The property Name is required.
type JsonTestParamsManySomethingThingy struct {
	// Name identifier
	Name string `json:"name,required"`
	// Optional count value
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

// Pet information
//
// The property Name is required.
type JsonTestParamsPet struct {
	// Name of the pet
	Name string `json:"name,required"`
	// Age of the pet in years
	Age param.Opt[int64] `json:"age,omitzero"`
	paramObj
}

func (r JsonTestParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow JsonTestParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTestParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User preference settings
type JsonTestParamsPreferences struct {
	// Whether to enable alert notifications
	Alerts param.Opt[bool] `json:"alerts,omitzero"`
	// UI theme preference (e.g., dark, light)
	Theme param.Opt[string] `json:"theme,omitzero"`
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

// An object with a required name and optional count
//
// The property Name is required.
type JsonTestParamsSomethingThingy struct {
	// Name identifier
	Name string `json:"name,required"`
	// Optional count value
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
	// A positive integer representing the new count value
	Body int64
	paramObj
}

func (r UpdateCountParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *UpdateCountParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type UploadTestParams struct {
	// The binary file to upload.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r UploadTestParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
