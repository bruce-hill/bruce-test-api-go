// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/DefinitelyATestOrg/test-api-go/internal/apijson"
	"github.com/DefinitelyATestOrg/test-api-go/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/test-api-go/option"
	"github.com/DefinitelyATestOrg/test-api-go/packages/pagination"
	"github.com/DefinitelyATestOrg/test-api-go/packages/param"
	"github.com/DefinitelyATestOrg/test-api-go/packages/respjson"
)

// PaginationService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaginationService] method instead.
type PaginationService struct {
	Options []option.RequestOption
	Ints    PaginationIntService
}

// NewPaginationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaginationService(opts ...option.RequestOption) (r PaginationService) {
	r = PaginationService{}
	r.Options = opts
	r.Ints = NewPaginationIntService(opts...)
	return
}

// Get foos
func (r *PaginationService) List(ctx context.Context, query PaginationListParams, opts ...option.RequestOption) (res *pagination.PageNumber[PaginationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "paginated"
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
func (r *PaginationService) ListAutoPaging(ctx context.Context, query PaginationListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[PaginationListResponse] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

type PaginationListResponse struct {
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
func (r PaginationListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginationListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	Tags []string         `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaginationListParams]'s query parameters as `url.Values`.
func (r PaginationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
