// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"github.com/bruce-hill/bruce-test-api-go/internal/apierror"
	"github.com/bruce-hill/bruce-test-api-go/packages/param"
	"github.com/bruce-hill/bruce-test-api-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AnimalUnionParam = shared.AnimalUnionParam

// This is an alias to an internal type.
type BirdParam = shared.BirdParam

// This is an alias to an internal type.
type CatParam = shared.CatParam

// This is an alias to an internal type.
type DogParam = shared.DogParam
