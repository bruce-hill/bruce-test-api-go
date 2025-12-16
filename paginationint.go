// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/DefinitelyATestOrg/test-api-go/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/test-api-go/option"
	"github.com/DefinitelyATestOrg/test-api-go/packages/pagination"
	"github.com/DefinitelyATestOrg/test-api-go/packages/param"
)

// PaginationIntService contains methods and other services that help with
// interacting with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaginationIntService] method instead.
type PaginationIntService struct {
	Options []option.RequestOption
}

// NewPaginationIntService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaginationIntService(opts ...option.RequestOption) (r PaginationIntService) {
	r = PaginationIntService{}
	r.Options = opts
	return
}

// Get paginated integers
func (r *PaginationIntService) List(ctx context.Context, query PaginationIntListParams, opts ...option.RequestOption) (res *pagination.PageNumber[int64], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "paginated-int"
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

// Get paginated integers
func (r *PaginationIntService) ListAutoPaging(ctx context.Context, query PaginationIntListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[int64] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

type PaginationIntListParams struct {
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaginationIntListParams]'s query parameters as
// `url.Values`.
func (r PaginationIntListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
