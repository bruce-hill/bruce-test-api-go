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

// TextService contains methods and other services that help with interacting with
// the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextService] method instead.
type TextService struct {
	Options []option.RequestOption
}

// NewTextService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTextService(opts ...option.RequestOption) (r TextService) {
	r = TextService{}
	r.Options = opts
	return
}

// Set the text that is returned when getting a Foo.
func (r *TextService) Set(ctx context.Context, body TextSetParams, opts ...option.RequestOption) (res *TextSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "foo-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type TextSetResponse = any

type TextSetParams struct {
	Name string `query:"name,required" json:"-"`
	paramObj
}

// URLQuery serializes [TextSetParams]'s query parameters as `url.Values`.
func (r TextSetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
