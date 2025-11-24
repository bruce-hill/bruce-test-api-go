// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/internal/requestconfig"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-go/packages/pagination"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
	"github.com/stainless-sdks/bruce-test-api-go/packages/respjson"
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

// Get foos
func (r *FooService) List(ctx context.Context, query FooListParams, opts ...option.RequestOption) (res *pagination.PageNumber[FooListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "foos"
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

// Get foos
func (r *FooService) ListAutoPaging(ctx context.Context, query FooListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[FooListResponse] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

type FooListResponse struct {
	// The baz field
	Baz int64 `json:"baz,required"`
	// The foo field
	Foo string `json:"foo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Baz         respjson.Field
		Foo         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
