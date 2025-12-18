// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/DefinitelyATestOrg/test-api-go/internal/apijson"
	shimjson "github.com/DefinitelyATestOrg/test-api-go/internal/encoding/json"
	"github.com/DefinitelyATestOrg/test-api-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/test-api-go/option"
	"github.com/DefinitelyATestOrg/test-api-go/packages/param"
	"github.com/DefinitelyATestOrg/test-api-go/packages/respjson"
)

// LinkedListService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkedListService] method instead.
type LinkedListService struct {
	Options []option.RequestOption
}

// NewLinkedListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLinkedListService(opts ...option.RequestOption) (r LinkedListService) {
	r = LinkedListService{}
	r.Options = opts
	return
}

// Accept a recursive linked list
func (r *LinkedListService) New(ctx context.Context, body LinkedListNewParams, opts ...option.RequestOption) (res *LinkedNode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "linked-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LinkedNode struct {
	Value string      `json:"value,required"`
	Next  *LinkedNode `json:"next,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedNode) RawJSON() string { return r.JSON.raw }
func (r *LinkedNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LinkedNode to a LinkedNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LinkedNodeParam.Overrides()
func (r LinkedNode) ToParam() LinkedNodeParam {
	return param.Override[LinkedNodeParam](json.RawMessage(r.RawJSON()))
}

// The property Value is required.
type LinkedNodeParam struct {
	Value string          `json:"value,required"`
	Next  LinkedNodeParam `json:"next,omitzero"`
	paramObj
}

func (r LinkedNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedListNewParams struct {
	LinkedNode LinkedNodeParam
	paramObj
}

func (r LinkedListNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LinkedNode)
}
func (r *LinkedListNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.LinkedNode)
}
