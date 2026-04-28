// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/bruce-hill/bruce-test-api-go/v2/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/v2/option"
)

// DoopService contains methods and other services that help with interacting with
// the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDoopService] method instead.
type DoopService struct {
	Options []option.RequestOption
}

// NewDoopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDoopService(opts ...option.RequestOption) (r DoopService) {
	r = DoopService{}
	r.Options = opts
	return
}

// Download a file using application/octet-stream
func (r *DoopService) Download(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "doop"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
