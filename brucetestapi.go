// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"bytes"
	"mime/multipart"
	"net/url"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apiform"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
)

type FooParams struct {
	Version     int64                `path:"version,required" json:"-"`
	Limit       param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Blorp       param.Opt[string]    `json:"blorp,omitzero"`
	XTraceID    param.Opt[string]    `header:"X-Trace-ID,omitzero" json:"-"`
	Filter      FooParamsFilter      `query:"filter,omitzero" json:"-"`
	Tags        []string             `query:"tags,omitzero" json:"-"`
	Preferences FooParamsPreferences `json:"preferences,omitzero"`
	XFlags      []string             `header:"X-Flags,omitzero" json:"-"`
	paramObj
}

func (r FooParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [FooParams]'s query parameters as `url.Values`.
func (r FooParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FooParamsFilter struct {
	Status param.Opt[string]   `query:"status,omitzero" json:"-"`
	Meta   FooParamsFilterMeta `query:"meta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FooParamsFilter]'s query parameters as `url.Values`.
func (r FooParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FooParamsFilterMeta struct {
	Level param.Opt[int64] `query:"level,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FooParamsFilterMeta]'s query parameters as `url.Values`.
func (r FooParamsFilterMeta) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FooParamsPreferences struct {
	Alerts param.Opt[bool]   `json:"alerts,omitzero"`
	Theme  param.Opt[string] `json:"theme,omitzero"`
	paramObj
}

func (r FooParamsPreferences) MarshalJSON() (data []byte, err error) {
	type shadow FooParamsPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FooParamsPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
