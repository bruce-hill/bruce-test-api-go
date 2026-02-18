// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/bruce-hill/bruce-test-api-go/internal/apijson"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

func AnimalParamOfBird(wingSpan float64) AnimalUnionParam {
	var variant BirdParam
	variant.WingSpan = wingSpan
	return AnimalUnionParam{OfBird: &variant}
}

func AnimalParamOfCat(lives int64, name string) AnimalUnionParam {
	var variant CatParam
	variant.Lives = lives
	variant.Name = name
	return AnimalUnionParam{OfCat: &variant}
}

func AnimalParamOfDog(goodBoy bool, name string) AnimalUnionParam {
	var variant DogParam
	variant.GoodBoy = goodBoy
	variant.Name = name
	return AnimalUnionParam{OfDog: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AnimalUnionParam struct {
	OfBird   *BirdParam        `json:",omitzero,inline"`
	OfCat    *CatParam         `json:",omitzero,inline"`
	OfDog    *DogParam         `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u AnimalUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBird,
		u.OfCat,
		u.OfDog,
		u.OfString,
		u.OfInt)
}
func (u *AnimalUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AnimalUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBird) {
		return u.OfBird
	} else if !param.IsOmitted(u.OfCat) {
		return u.OfCat
	} else if !param.IsOmitted(u.OfDog) {
		return u.OfDog
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnimalUnionParam) GetWingSpan() *float64 {
	if vt := u.OfBird; vt != nil {
		return &vt.WingSpan
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnimalUnionParam) GetLives() *int64 {
	if vt := u.OfCat; vt != nil {
		return &vt.Lives
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnimalUnionParam) GetGoodBoy() *bool {
	if vt := u.OfDog; vt != nil {
		return &vt.GoodBoy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnimalUnionParam) GetType() *string {
	if vt := u.OfBird; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCat; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDog; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AnimalUnionParam) GetName() *string {
	if vt := u.OfCat; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDog; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// The properties Type, WingSpan are required.
type BirdParam struct {
	WingSpan float64 `json:"wingSpan,required"`
	// This field can be elided, and will marshal its zero value as "bird".
	Type constant.Bird `json:"type,required"`
	paramObj
}

func (r BirdParam) MarshalJSON() (data []byte, err error) {
	type shadow BirdParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BirdParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lives, Name, Type are required.
type CatParam struct {
	Lives int64  `json:"lives,required"`
	Name  string `json:"name,required"`
	// This field can be elided, and will marshal its zero value as "cat".
	Type constant.Cat `json:"type,required"`
	paramObj
}

func (r CatParam) MarshalJSON() (data []byte, err error) {
	type shadow CatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties GoodBoy, Name, Type are required.
type DogParam struct {
	GoodBoy bool   `json:"goodBoy,required"`
	Name    string `json:"name,required"`
	// This field can be elided, and will marshal its zero value as "dog".
	Type constant.Dog `json:"type,required"`
	paramObj
}

func (r DogParam) MarshalJSON() (data []byte, err error) {
	type shadow DogParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DogParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
