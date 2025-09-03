// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
)

// NameService contains methods and other services that help with interacting with
// the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNameService] method instead.
type NameService struct {
	Options []option.RequestOption
}

// NewNameService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNameService(opts ...option.RequestOption) (r NameService) {
	r = NameService{}
	r.Options = opts
	return
}

// Set Name
func (r *NameService) Set(ctx context.Context, body NameSetParams, opts ...option.RequestOption) (res *NameSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "name"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type NameSetResponse = any

type NameSetParams struct {
	Name string `query:"name,required" json:"-"`
	paramObj
}

// URLQuery serializes [NameSetParams]'s query parameters as `url.Values`.
func (r NameSetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
