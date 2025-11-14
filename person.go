// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/internal/requestconfig"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-go/packages/pagination"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
	"github.com/stainless-sdks/bruce-test-api-go/packages/respjson"
)

// PersonService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	Options []option.RequestOption
	Pets    PersonPetService
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.Options = opts
	r.Pets = NewPersonPetService(opts...)
	return
}

// Create a new person and add them to the system.
func (r *PersonService) New(ctx context.Context, body PersonNewParams, opts ...option.RequestOption) (res *PersonNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "people"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a person's information by ID.
func (r *PersonService) Get(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing person's information.
func (r *PersonService) Update(ctx context.Context, personID string, body PersonUpdateParams, opts ...option.RequestOption) (res *PersonUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a list of all people.
func (r *PersonService) List(ctx context.Context, query PersonListParams, opts ...option.RequestOption) (res *pagination.PageNumber[PersonListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "people"
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

// Get a list of all people.
func (r *PersonService) ListAutoPaging(ctx context.Context, query PersonListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[PersonListResponse] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

// Remove a person from the system.
func (r *PersonService) Delete(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PersonNewResponse struct {
	// The person's job
	Job string `json:"job,required"`
	// The person's name
	Name PersonNewResponseName `json:"name,required"`
	// Unique person identifier
	ID string `json:"id" format:"uuid7"`
	// Image of the person (base64)
	ImageBase64 string `json:"image_base64,nullable" format:"byte"`
	// Image of the person (binary)
	ImageBinary io.Reader `json:"image_binary,nullable" format:"binary"`
	// The person's pets
	Pets []PersonNewResponsePet `json:"pets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		ImageBase64 respjson.Field
		ImageBinary respjson.Field
		Pets        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person's name
type PersonNewResponseName struct {
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
func (r PersonNewResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonNewResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonNewResponsePet struct {
	// The pet's name
	Name PersonNewResponsePetName `json:"name,required"`
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
func (r PersonNewResponsePet) RawJSON() string { return r.JSON.raw }
func (r *PersonNewResponsePet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonNewResponsePetName struct {
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
func (r PersonNewResponsePetName) RawJSON() string { return r.JSON.raw }
func (r *PersonNewResponsePetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonGetResponse struct {
	// The person's job
	Job string `json:"job,required"`
	// The person's name
	Name PersonGetResponseName `json:"name,required"`
	// Unique person identifier
	ID string `json:"id" format:"uuid7"`
	// Image of the person (base64)
	ImageBase64 string `json:"image_base64,nullable" format:"byte"`
	// Image of the person (binary)
	ImageBinary io.Reader `json:"image_binary,nullable" format:"binary"`
	// The person's pets
	Pets []PersonGetResponsePet `json:"pets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		ImageBase64 respjson.Field
		ImageBinary respjson.Field
		Pets        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person's name
type PersonGetResponseName struct {
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
func (r PersonGetResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonGetResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonGetResponsePet struct {
	// The pet's name
	Name PersonGetResponsePetName `json:"name,required"`
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
func (r PersonGetResponsePet) RawJSON() string { return r.JSON.raw }
func (r *PersonGetResponsePet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonGetResponsePetName struct {
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
func (r PersonGetResponsePetName) RawJSON() string { return r.JSON.raw }
func (r *PersonGetResponsePetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonUpdateResponse struct {
	// The person's job
	Job string `json:"job,required"`
	// The person's name
	Name PersonUpdateResponseName `json:"name,required"`
	// Unique person identifier
	ID string `json:"id" format:"uuid7"`
	// Image of the person (base64)
	ImageBase64 string `json:"image_base64,nullable" format:"byte"`
	// Image of the person (binary)
	ImageBinary io.Reader `json:"image_binary,nullable" format:"binary"`
	// The person's pets
	Pets []PersonUpdateResponsePet `json:"pets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		ImageBase64 respjson.Field
		ImageBinary respjson.Field
		Pets        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person's name
type PersonUpdateResponseName struct {
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
func (r PersonUpdateResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonUpdateResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonUpdateResponsePet struct {
	// The pet's name
	Name PersonUpdateResponsePetName `json:"name,required"`
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
func (r PersonUpdateResponsePet) RawJSON() string { return r.JSON.raw }
func (r *PersonUpdateResponsePet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonUpdateResponsePetName struct {
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
func (r PersonUpdateResponsePetName) RawJSON() string { return r.JSON.raw }
func (r *PersonUpdateResponsePetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponse struct {
	// The person's job
	Job string `json:"job,required"`
	// The person's name
	Name PersonListResponseName `json:"name,required"`
	// Unique person identifier
	ID string `json:"id" format:"uuid7"`
	// Image of the person (base64)
	ImageBase64 string `json:"image_base64,nullable" format:"byte"`
	// Image of the person (binary)
	ImageBinary io.Reader `json:"image_binary,nullable" format:"binary"`
	// The person's pets
	Pets []PersonListResponsePet `json:"pets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		ImageBase64 respjson.Field
		ImageBinary respjson.Field
		Pets        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person's name
type PersonListResponseName struct {
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
func (r PersonListResponseName) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponsePet struct {
	// The pet's name
	Name PersonListResponsePetName `json:"name,required"`
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
func (r PersonListResponsePet) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponsePet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
type PersonListResponsePetName struct {
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
func (r PersonListResponsePetName) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponsePetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonDeleteResponse map[string]any

type PersonNewParams struct {
	// The name of the person to create
	Name PersonNewParamsName `json:"name,omitzero,required"`
	// The person's job
	Job param.Opt[string] `json:"job,omitzero"`
	// A list of pets for this person
	Pets []PersonNewParamsPet `json:"pets,omitzero"`
	paramObj
}

func (r PersonNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the person to create
//
// The property FullName is required.
type PersonNewParamsName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PersonNewParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Species are required.
type PersonNewParamsPet struct {
	// The pet's name
	Name PersonNewParamsPetName `json:"name,omitzero,required"`
	// The pet's species
	Species string `json:"species,required"`
	paramObj
}

func (r PersonNewParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
//
// The property FullName is required.
type PersonNewParamsPetName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PersonNewParamsPetName) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParamsPetName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParamsPetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonUpdateParams struct {
	// The updated name of the person
	Name PersonUpdateParamsName `json:"name,omitzero,required"`
	// The updated job of the person
	Job param.Opt[string] `json:"job,omitzero"`
	paramObj
}

func (r PersonUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The updated name of the person
//
// The property FullName is required.
type PersonUpdateParamsName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PersonUpdateParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonUpdateParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonUpdateParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListParams struct {
	// Job name to search for
	Job param.Opt[string] `query:"job,omitzero" json:"-"`
	// Full name to search for
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Nickname to search for
	Nickname param.Opt[string] `query:"nickname,omitzero" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
