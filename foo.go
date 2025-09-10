// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/bruce-hill/bruce-test-api-go/packages/pagination"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
)

// FooService contains methods and other services that help with interacting with
// the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFooService] method instead.
type FooService struct {
	Options []option.RequestOption
}

// NewFooService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFooService(opts ...option.RequestOption) (r FooService) {
	r = FooService{}
	r.Options = opts
	return
}

// Get a Foo that has text, a random number, and a list of random numbers.
func (r *FooService) Get(ctx context.Context, opts ...option.RequestOption) (res *FooGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "foo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all of the Foos.
func (r *FooService) List(ctx context.Context, query FooListParams, opts ...option.RequestOption) (res *pagination.PageNumber[FooListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "all-foos"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of all of the Foos.
func (r *FooService) ListAutoPaging(ctx context.Context, query FooListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[FooListResponse] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

type FooGetResponse struct {
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
func (r FooGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FooGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FooListResponse struct {
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
func (r FooListResponse) RawJSON() string { return r.JSON.raw }
func (r *FooListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FooListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FooListParams]'s query parameters as `url.Values`.
func (r FooListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
