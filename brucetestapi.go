// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/DefinitelyATestOrg/test-api-go/internal/apiform"
	"github.com/DefinitelyATestOrg/test-api-go/internal/apijson"
	"github.com/DefinitelyATestOrg/test-api-go/internal/apiquery"
	shimjson "github.com/DefinitelyATestOrg/test-api-go/internal/encoding/json"
	"github.com/DefinitelyATestOrg/test-api-go/packages/param"
)

type FormTestParams struct {
	Version     int64                     `path:"version,required" json:"-"`
	Date        time.Time                 `query:"date,required" format:"date" json:"-"`
	Datetime    time.Time                 `query:"datetime,required" format:"date-time" json:"-"`
	Time        string                    `query:"time,required" format:"time" json:"-"`
	Limit       param.Opt[int64]          `query:"limit,omitzero" json:"-"`
	Blorp       param.Opt[string]         `json:"blorp,omitzero"`
	XTraceID    param.Opt[string]         `header:"X-Trace-ID,omitzero" json:"-"`
	PlsNull     any                       `json:"pls_null"`
	Filter      FormTestParamsFilter      `query:"filter,omitzero" json:"-"`
	Tags        []string                  `query:"tags,omitzero" json:"-"`
	Preferences FormTestParamsPreferences `json:"preferences,omitzero"`
	XFlags      []string                  `header:"X-Flags,omitzero" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
	Version     int64                     `path:"version,required" json:"-"`
	Date        time.Time                 `query:"date,required" format:"date" json:"-"`
	Datetime    time.Time                 `query:"datetime,required" format:"date-time" json:"-"`
	Time        string                    `query:"time,required" format:"time" json:"-"`
	Limit       param.Opt[int64]          `query:"limit,omitzero" json:"-"`
	Blorp       param.Opt[string]         `json:"blorp,omitzero"`
	XTraceID    param.Opt[string]         `header:"X-Trace-ID,omitzero" json:"-"`
	PlsNull     any                       `json:"pls_null"`
	Filter      JsonTestParamsFilter      `query:"filter,omitzero" json:"-"`
	Tags        []string                  `query:"tags,omitzero" json:"-"`
	Preferences JsonTestParamsPreferences `json:"preferences,omitzero"`
	XFlags      []string                  `header:"X-Flags,omitzero" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
