// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/bruce-hill/stainless-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/stainless-test-api-go/option"
	"github.com/bruce-hill/stainless-test-api-go/packages/jsonl"
)

// StreamJsonService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamJsonService] method instead.
type StreamJsonService struct {
	Options []option.RequestOption
}

// NewStreamJsonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamJsonService(opts ...option.RequestOption) (r StreamJsonService) {
	r = StreamJsonService{}
	r.Options = opts
	return
}

// Streams JSON objects as a chunked response (NDJSON)
func (r *StreamJsonService) StreamStreaming(ctx context.Context, opts ...option.RequestOption) (stream *jsonl.Stream[StreamJsonStreamResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-ndjson")}, opts...)
	path := "stream-json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return jsonl.NewStream[StreamJsonStreamResponse](raw, err)
}

type StreamJsonStreamResponse map[string]any
