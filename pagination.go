// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
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

// Retrieves a paginated list of Foo objects with optional filtering by tags.
// Supports standard pagination parameters.
func (r *PaginationService) List(ctx context.Context, query PaginationListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "paginated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type PaginationListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter results by tags
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaginationListParams]'s query parameters as `url.Values`.
func (r PaginationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
