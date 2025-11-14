// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apiform"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
)

type FnordResponse map[string]any

type PostFnordResponse map[string]any

type TestFormResponse = any

type FnordParams struct {
	// The first positional arg
	FirstPos string `path:"first_pos,required" json:"-"`
	// The first query param (required)
	FirstQuery []int64 `query:"first_query,omitzero,required" json:"-"`
	// The second query param (optional)
	SecondQuery param.Opt[string] `query:"second_query,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FnordParams]'s query parameters as `url.Values`.
func (r FnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PostFnordParams struct {
	// The first positional arg
	FirstPos string `path:"first_pos,required" json:"-"`
	// The first query param (required)
	ArrayItems []int64 `query:"array_items,omitzero,required" json:"-"`
	// The name of the person to create
	Name PostFnordParamsName `json:"name,omitzero,required"`
	// The second query param (optional)
	SecondQuery param.Opt[string] `query:"second_query,omitzero" json:"-"`
	// Image of the person (base64)
	ImageBase64 param.Opt[string] `json:"image_base64,omitzero" format:"byte"`
	// The person's job
	Job param.Opt[string] `json:"job,omitzero"`
	// Image of the person (binary)
	ImageBinary io.Reader `json:"image_binary,omitzero" format:"binary"`
	// A list of pets for this person
	Pets []PostFnordParamsPet `json:"pets,omitzero"`
	paramObj
}

func (r PostFnordParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [PostFnordParams]'s query parameters as `url.Values`.
func (r PostFnordParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The name of the person to create
//
// The property FullName is required.
type PostFnordParamsName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PostFnordParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PostFnordParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostFnordParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Species are required.
type PostFnordParamsPet struct {
	// The pet's name
	Name PostFnordParamsPetName `json:"name,omitzero,required"`
	// The pet's species
	Species string `json:"species,required"`
	paramObj
}

func (r PostFnordParamsPet) MarshalJSON() (data []byte, err error) {
	type shadow PostFnordParamsPet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostFnordParamsPet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pet's name
//
// The property FullName is required.
type PostFnordParamsPetName struct {
	// Full name
	FullName string `json:"full_name,required"`
	// Nickname (if different from full name)
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	paramObj
}

func (r PostFnordParamsPetName) MarshalJSON() (data []byte, err error) {
	type shadow PostFnordParamsPetName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostFnordParamsPetName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestFormParams struct {
	// Email
	Email string `json:"email,required"`
	// Username
	Username string `json:"username,required" format:"byte"`
	// Age
	Age param.Opt[int64] `json:"age,omitzero"`
	// Subscribe
	Subscribe param.Opt[bool] `json:"subscribe,omitzero"`
	paramObj
}

func (r TestFormParams) MarshalJSON() (data []byte, err error) {
	type shadow TestFormParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestFormParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
