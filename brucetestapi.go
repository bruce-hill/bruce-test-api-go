// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"net/url"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
)

type FnordResponse map[string]string

type FnordParams struct {
	// The first positional arg
	FirstPos string `path:"first_pos,required" json:"-"`
	// The first query param (required)
	FirstQuery string `query:"first_query,required" json:"-"`
	// The second query param (optional)
	SecondQuery param.Opt[string] `query:"second_query,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FnordParams]'s query parameters as `url.Values`.
func (r FnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
