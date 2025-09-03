// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"net/url"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
)

type GetFooResponse struct {
	ListOfNums   []int64 `json:"list_of_nums,required"`
	RandomNumber int64   `json:"random_number,required"`
	Text         string  `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListOfNums   respjson.Field
		RandomNumber respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetFooResponse) RawJSON() string { return r.JSON.raw }
func (r *GetFooResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetTextResponse = any

type SetTextParams struct {
	Name string `query:"name,required" json:"-"`
	paramObj
}

// URLQuery serializes [SetTextParams]'s query parameters as `url.Values`.
func (r SetTextParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
