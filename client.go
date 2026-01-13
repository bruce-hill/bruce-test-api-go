// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"

	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
)

// Client creates a struct with services and top level methods that help with
// interacting with the bruce-test-api API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options    []option.RequestOption
	Pagination PaginationService
	StreamJson StreamJsonService
}

// DefaultClientOptions read from the environment (BRUCE_TEST_API_API_KEY,
// BRUCE_TEST_API_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BRUCE_TEST_API_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("BRUCE_TEST_API_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (BRUCE_TEST_API_API_KEY, BRUCE_TEST_API_BASE_URL). The option passed
// in as arguments are applied after these default arguments, and all option will
// be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Pagination = NewPaginationService(opts...)
	r.StreamJson = NewStreamJsonService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}

// Demonstrates a form-data endpoint with various parameter types including path,
// query, and header parameters. Accepts multipart form data for user updates.
func (r *Client) FormTest(ctx context.Context, userID string, params FormTestParams, opts ...option.RequestOption) (res *FormTestResponse, err error) {
	for _, v := range params.XFlags {
		opts = append(opts, option.WithHeaderAdd("X-Flags", fmt.Sprintf("%s", v)))
	}
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-ID", fmt.Sprintf("%s", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("form-v%v/users/%s", params.Version, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Demonstrates a JSON endpoint with various parameter types including path, query,
// and header parameters. Accepts JSON body for user updates.
func (r *Client) JsonTest(ctx context.Context, userID string, params JsonTestParams, opts ...option.RequestOption) (res *JsonTestResponse, err error) {
	for _, v := range params.XFlags {
		opts = append(opts, option.WithHeaderAdd("X-Flags", fmt.Sprintf("%s", v)))
	}
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-ID", fmt.Sprintf("%s", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("json-v%v/users/%s", params.Version, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates the current count with a new integer value. The value must be a positive
// integer (minimum 1).
func (r *Client) UpdateCount(ctx context.Context, body UpdateCountParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Accepts a binary file upload using multipart/form-data. Typical use cases
// include uploading images, documents, or other opaque binaries.
func (r *Client) UploadTest(ctx context.Context, body UploadTestParams, opts ...option.RequestOption) (res *UploadTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
