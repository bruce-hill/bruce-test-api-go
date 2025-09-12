// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go/internal/requestconfig"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/bruce-hill/bruce-test-api-go/packages/pagination"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
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
func (r *PersonService) New(ctx context.Context, body PersonNewParams, opts ...option.RequestOption) (res *Person, err error) {
	opts = append(r.Options[:], opts...)
	path := "people"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a person's information by ID.
func (r *PersonService) Get(ctx context.Context, personID string, opts ...option.RequestOption) (res *Person, err error) {
	opts = append(r.Options[:], opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing person's information.
func (r *PersonService) Update(ctx context.Context, personID string, body PersonUpdateParams, opts ...option.RequestOption) (res *Person, err error) {
	opts = append(r.Options[:], opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a list of all people.
func (r *PersonService) List(ctx context.Context, query PersonListParams, opts ...option.RequestOption) (res *pagination.PageNumber[Person], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *PersonService) ListAutoPaging(ctx context.Context, query PersonListParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[Person] {
	return pagination.NewPageNumberAutoPager(r.List(ctx, query, opts...))
}

// Remove a person from the system.
func (r *PersonService) Delete(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Name struct {
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
func (r Name) RawJSON() string { return r.JSON.raw }
func (r *Name) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Name to a NameParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NameParam.Overrides()
func (r Name) ToParam() NameParam {
	return param.Override[NameParam](json.RawMessage(r.RawJSON()))
}

// The property FullName is required.
type NameParam struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r NameParam) MarshalJSON() (data []byte, err error) {
	type shadow NameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Person struct {
	// The person's job
	Job string `json:"job,required"`
	// The person's name
	Name Name `json:"name,required"`
	// Unique person identifier
	ID string `json:"id" format:"uuid7"`
	// The person's pets
	Pets []Pet `json:"pets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Pets        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Person) RawJSON() string { return r.JSON.raw }
func (r *Person) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonDeleteResponse map[string]any

type PersonNewParams struct {
	// The name of the person to create
	Name NameParam `json:"name,omitzero,required"`
	// The person's job
	Job param.Opt[string] `json:"job,omitzero"`
	// A list of pet names to create as pets for this person
	PetNames []NameParam `json:"pet_names,omitzero"`
	paramObj
}

func (r PersonNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonUpdateParams struct {
	// The updated name of the person
	Name NameParam `json:"name,omitzero,required"`
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

type PersonListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
