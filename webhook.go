// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/requestconfig"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
	"github.com/stainless-sdks/bruce-test-api-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Register Webhook
func (r *WebhookService) Register(ctx context.Context, body WebhookRegisterParams, opts ...option.RequestOption) (res *WebhookRegisterResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "register-webhook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *WebhookService) Unwrap(body []byte) (*UnwrapWebhookEventUnion, error) {
	res := &UnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(body)
	if err != nil {
		return res, err
	}
	return res, nil
}

type WebhookRegisterResponse struct {
	Secret string `json:"secret,required"`
	URL    string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThingHappenedWebhookEvent struct {
	Thing string `json:"thing,required"`
	// Any of "thingHappened".
	Type ThingHappenedWebhookEventType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thing       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThingHappenedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ThingHappenedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThingHappenedWebhookEventType string

const (
	ThingHappenedWebhookEventTypeThingHappened ThingHappenedWebhookEventType = "thingHappened"
)

type ThingDidntHappenWebhookEvent struct {
	Thing string `json:"thing,required"`
	// Any of "thingDidntHappen".
	Type ThingDidntHappenWebhookEventType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thing       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThingDidntHappenWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ThingDidntHappenWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThingDidntHappenWebhookEventType string

const (
	ThingDidntHappenWebhookEventTypeThingDidntHappen ThingDidntHappenWebhookEventType = "thingDidntHappen"
)

// UnwrapWebhookEventUnion contains all possible properties and values from
// [ThingHappenedWebhookEvent], [ThingDidntHappenWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	Thing string `json:"thing"`
	Type  string `json:"type"`
	JSON  struct {
		Thing respjson.Field
		Type  respjson.Field
		raw   string
	} `json:"-"`
}

func (u UnwrapWebhookEventUnion) AsThingHappened() (v ThingHappenedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsThingDidntHappen() (v ThingDidntHappenWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookRegisterParams struct {
	URL string `json:"url,required" format:"uri"`
	paramObj
}

func (r WebhookRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
