// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
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

// Foo
func (r *FooService) List(ctx context.Context, opts ...option.RequestOption) (res *FooListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "foo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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
