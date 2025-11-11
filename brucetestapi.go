// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"net/url"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
)

type FnordResponse map[string]any

type PostFnordResponse map[string]any

type FnordParams struct {
	// The first positional arg
	FirstPos string `path:"first_pos,required" json:"-"`
	// The first query param (required)
	ArrayItems []int64 `query:"array_items,omitzero,required" json:"-"`
	// The second query param (optional)
	SecondQuery FnordParamsSecondQuery `query:"second_query,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FnordParams]'s query parameters as `url.Values`.
func (r FnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The second query param (optional)
type FnordParamsSecondQuery struct {
	Age  param.Opt[int64]  `query:"age,omitzero" json:"-"`
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FnordParamsSecondQuery]'s query parameters as `url.Values`.
func (r FnordParamsSecondQuery) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PostFnordParams struct {
	// The first positional arg
	FirstPos string `path:"first_pos,required" json:"-"`
	// The first query param (required)
	FirstQuery []int64 `query:"first_query,omitzero,required" json:"-"`
	// Full name
	FullName string `json:"full_name,required"`
	// The second query param (optional)
	SecondQuery param.Opt[string] `query:"second_query,omitzero" json:"-"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PostFnordParams) MarshalJSON() (data []byte, err error) {
	type shadow PostFnordParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostFnordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PostFnordParams]'s query parameters as `url.Values`.
func (r PostFnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
