// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/internal/requestconfig"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
	"github.com/stainless-sdks/bruce-test-api-go/packages/respjson"
)

// PersonPetService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonPetService] method instead.
type PersonPetService struct {
	Options []option.RequestOption
}

// NewPersonPetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPersonPetService(opts ...option.RequestOption) (r PersonPetService) {
	r = PersonPetService{}
	r.Options = opts
	return
}

// Add a new pet to an existing person.
func (r *PersonPetService) New(ctx context.Context, personID string, body PersonPetNewParams, opts ...option.RequestOption) (res *PersonPetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s/pets", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a pet from a person.
func (r *PersonPetService) Get(ctx context.Context, personID string, query PersonPetGetParams, opts ...option.RequestOption) (res *PersonPetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s/pet", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update an existing pet's information.
func (r *PersonPetService) Update(ctx context.Context, petID string, params PersonPetUpdateParams, opts ...option.RequestOption) (res *PersonPetUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.PersonID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	if petID == "" {
		err = errors.New("missing required pet_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s/pets/%s", params.PersonID, petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get all pets belonging to a specific person by their ID.
func (r *PersonPetService) List(ctx context.Context, personID string, opts ...option.RequestOption) (res *[]PersonPetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s/pets", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a pet from a person.
func (r *PersonPetService) Delete(ctx context.Context, petID string, body PersonPetDeleteParams, opts ...option.RequestOption) (res *PersonPetDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PersonID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	if petID == "" {
		err = errors.New("missing required pet_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s/pets/%s", body.PersonID, petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PersonPetNewResponse struct {
	// The pet's name
	Name PersonPetNewResponseName `json:"name,required"`
	// The pet's species
	Species string `json:"species,required"`
	// Unique pet identifier
	ID string `json:"id" format:"uuid7"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Species     respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonPetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonPetNewResponseName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname string `json:"nickname,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Nickname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetNewResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonPetNewResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetGetResponse struct {
	// The pet's name
	Name PersonPetGetResponseName `json:"name,required"`
	// The pet's species
	Species string `json:"species,required"`
	// Unique pet identifier
	ID string `json:"id" format:"uuid7"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Species     respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonPetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonPetGetResponseName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname string `json:"nickname,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Nickname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetGetResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonPetGetResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetUpdateResponse struct {
	// The pet's name
	Name PersonPetUpdateResponseName `json:"name,required"`
	// The pet's species
	Species string `json:"species,required"`
	// Unique pet identifier
	ID string `json:"id" format:"uuid7"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Species     respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonPetUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonPetUpdateResponseName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname string `json:"nickname,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Nickname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetUpdateResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonPetUpdateResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetListResponse struct {
	// The pet's name
	Name PersonPetListResponseName `json:"name,required"`
	// The pet's species
	Species string `json:"species,required"`
	// Unique pet identifier
	ID string `json:"id" format:"uuid7"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Species     respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonPetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonPetListResponseName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname string `json:"nickname,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName    respjson.Field
		Nickname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonPetListResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonPetListResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetDeleteResponse map[string]any

type PersonPetNewParams struct {
	// The name of the pet to create
	Name PersonPetNewParamsName `json:"name,omitzero,required"`
	// The species of the pet
	Species param.Opt[string] `json:"species,omitzero"`
	paramObj
}

func (r PersonPetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonPetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonPetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the pet to create
//
// The property FullName is required.
type PersonPetNewParamsName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PersonPetNewParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonPetNewParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonPetNewParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetGetParams struct {
	// The pet's name
	PetName string `query:"pet_name,required" json:"-"`
	paramObj
}

// URLQuery serializes [PersonPetGetParams]'s query parameters as `url.Values`.
func (r PersonPetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonPetUpdateParams struct {
	// The unique identifier of the person who owns the pet
	PersonID string `path:"person_id,required" format:"uuid" json:"-"`
	// The updated name of the pet
	Name PersonPetUpdateParamsName `json:"name,omitzero,required"`
	// The updated species of the pet
	Species string `json:"species,required"`
	paramObj
}

func (r PersonPetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonPetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonPetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The updated name of the pet
//
// The property FullName is required.
type PersonPetUpdateParamsName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PersonPetUpdateParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonPetUpdateParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonPetUpdateParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonPetDeleteParams struct {
	// The unique identifier of the person who owns the pet
	PersonID string `path:"person_id,required" format:"uuid" json:"-"`
	paramObj
}
