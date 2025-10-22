// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"net/url"

	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
)

type FnordResponse map[string]string

type FnordParams struct {
	// The first positional arg
	Pos1 string `path:"pos1,required" json:"-"`
	// The first query param (required)
	Query1 string `query:"query1,required" json:"-"`
	// The second query param (optional)
	Query2 param.Opt[string] `query:"query2,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FnordParams]'s query parameters as `url.Values`.
func (r FnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
