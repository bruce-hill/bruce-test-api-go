// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"encoding/json"
	"time"

	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	shimjson "github.com/bruce-hill/bruce-test-api-go/internal/encoding/json"
	"github.com/bruce-hill/bruce-test-api-go/packages/respjson"
	"github.com/bruce-hill/bruce-test-api-go/shared"
)

// Response confirming the user update
type NewAnimalResponse struct {
	// Status of the update operation
	Status string `json:"status"`
	// Timestamp when the update occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// The updated user's ID
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewAnimalResponse) RawJSON() string { return r.JSON.raw }
func (r *NewAnimalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewAnimalParams struct {
	Body shared.AnimalUnionParam
	paramObj
}

func (r NewAnimalParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *NewAnimalParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
